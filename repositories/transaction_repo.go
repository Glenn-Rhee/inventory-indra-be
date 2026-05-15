package repositories

import (
	"errors"
	"fmt"
	"inventory-indra/model"
	"log"
	"math"
	"net/http"
	"time"

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

	createdAt := time.Now()
	var idTransaction = createdAt.Format("TX-02012006-150405")

	transaction := model.Transaction{
		Id: idTransaction,
		ProductId: product.Id,
		TransactionType: data.TransactionType,
		Quantity: data.Quantity,
		CreatedAt: createdAt,
	}

	result = tx.Model(&transaction).Create(transaction)
	if result.Error != nil {
		tx.Rollback()
		log.Println("Error while create data transaction:", result.Error.Error())
		return errors.New("An error while create transaction! Please try again later"), http.StatusInternalServerError
	}

	if data.TransactionType == model.TransactionIN && data.Quantity > product.Stock.StockPerButir {
		tx.Rollback()
		return fmt.Errorf("Out of stock! Maximum stock is %d", product.Stock.StockPerButir), http.StatusBadRequest
	}

	var stockLeft int
	if data.TransactionType == model.TransactionIN {
		stockLeft = product.Stock.StockPerButir - data.Quantity
	} else {
		stockLeft = product.Stock.StockPerButir + data.Quantity
	}


	result = tx.Model(&model.Stock{}).Where("product_id = ?", product.Id).Updates(model.Stock{
		StockPerButir: stockLeft,
	})
	if result.Error != nil {
		tx.Rollback()
		log.Println("Error while update stock:", result.Error.Error())
		return errors.New("An error while create transaction! Please try again later"), http.StatusInternalServerError
	}

	if data.TransactionType == model.TransactionOUT {
		if data.ExpiredDate == nil {
			tx.Rollback()
			return errors.New("Expired date is required!"), http.StatusBadRequest
		}

		result = tx.Model(&model.Product{}).Where("id = ?", product.Id).Updates(model.Product{
			ExpiredDate: *data.ExpiredDate,
			PricePerButir: data.Price,
		})

		if result.Error != nil {
			tx.Rollback()
			log.Println("Error while update expired date:", result.Error.Error())
			return errors.New("An error while create transaction! Please try again later"), http.StatusInternalServerError
		}
	}

	tx.Commit()
	return nil, http.StatusCreated
}

type GetTransactionParams struct {
	Limit 	int
	Page 	int
	Filter 	string
}

func (r *TransactionRepository) GetTransaction(params GetTransactionParams) (model.GetTransaction, error, int) {
	var transactions []model.Transaction
	var totalRows int64
	offset := (params.Page - 1) * params.Limit
	query := r.db.Model(&model.Transaction{}).
		Joins("JOIN products ON products.id = transactions.product_id").
		Preload("Product")

	if params.Filter != "" {
		query = query.Where(
			"transactions.id = ? OR products.name ILIKE ? OR transactions.transaction_type::text ILIKE ?", 
			params.Filter, 
			"%" + params.Filter + "%",
			"%" + params.Filter + "%",
		)
	}

	query.Count(&totalRows)
	
	result := query.Order("transactions.created_at DESC").
				Limit(params.Limit).
				Offset(offset).
				Find(&transactions)
	
	if result.Error != nil {
		log.Println("Error while get data transaction:", result.Error.Error())
		return model.GetTransaction{}, errors.New("An error while get data transaction! Please try again later!"), http.StatusInternalServerError
	}

	transactionsResponse := make([]model.TransactionResponse, len(transactions))
	totalPages := int(math.Ceil(float64(totalRows) / float64(params.Limit)))
	var totalRevenue = 0
	var totalPurchase = 0
	for i, transaction := range transactions {
		totalPrice := transaction.Quantity * transaction.Product.PricePerButir
		if transaction.TransactionType == model.TransactionIN {
			totalRevenue += totalPrice
		} else {
			totalPurchase += totalPrice
		}

		transactionsResponse[i] = model.TransactionResponse{
			Id: transaction.Id,
			ProductName: transaction.Product.Name,
			TransactionType: transaction.TransactionType,
			Quantity: transaction.Quantity,
			Price:int64(transaction.Product.PricePerButir),
			TotalPrice: int64(totalPrice),
			TransactionDate: transaction.CreatedAt,
		}
	}

	return model.GetTransaction{
		TotalTransaction: int(totalRows),
		TotalRevenue: totalRevenue,
		TotalPurchase: totalPurchase,
		Transactions: transactionsResponse,
		TotalPages: totalPages,
	}, nil, http.StatusOK
}