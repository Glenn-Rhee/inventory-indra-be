package repositories

import (
	"errors"
	"fmt"
	"inventory-indra/model"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateTransaction(data model.CreateTransaction) (error, int) {
	var product model.Product

	result := r.db.Where("products.id = ?", data.ProductId).Joins("Stock").First(&product)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println("Error find product:", result.Error.Error())
		return errors.New("Product not found!"), http.StatusNotFound
	}

	tx := r.db.Begin()
	if tx.Error != nil {
		log.Println("Error setup transaction:", tx.Error.Error())
		return errors.New("An error while create transaction! Please try again later"), http.StatusInternalServerError
	}

	transaction := model.Transaction{
		Id: uuid.New().String(),
		ProductId: product.Id,
		TransactionType: data.TransactionType,
		Quantity: data.Quantity,
		CreatedAt: time.Now(),
	}

	result = tx.Model(&transaction)
	if result.Error != nil {
		tx.Rollback()
		log.Println("Error while create data transaction:", result.Error.Error())
		return errors.New("An error while create transaction! Please try again later"), http.StatusInternalServerError
	}

	if data.Quantity > product.Stock.StockPerButir {
		tx.Rollback()
		return fmt.Errorf("Out of stock! Maximum stock is %d", product.Stock.StockPerButir), http.StatusBadRequest
	}

	stockLeft := product.Stock.StockPerButir - data.Quantity

	result = tx.Model(&model.Stock{}).Where("product_id = ?", product.Id).Updates(model.Stock{
		StockPerButir: stockLeft,
	})
	if result.Error != nil {
		tx.Rollback()
		log.Println("Error while update stock:", result.Error.Error())
		return errors.New("An error while create transaction! Please try again later"), http.StatusInternalServerError
	}

	tx.Commit()
	return nil, http.StatusCreated
}