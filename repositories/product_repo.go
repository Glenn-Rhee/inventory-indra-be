package repositories

import (
	"errors"
	"inventory-indra/helper"
	"inventory-indra/model"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}



func NewProductRepository(db *gorm.DB) *ProductRepository{
	return  &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(dataProduct model.CreateProduct) (error, int){
	var product model.Product

	result := r.db.Where("name = ?", dataProduct.Name).First(&product)
	tx := r.db.Begin()

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {

		if tx.Error != nil {
			return errors.New("Failed to start transaction! Please try again later"), http.StatusInternalServerError
		}
		// Update product
		result = tx.Model(&model.Product{}).Where("id = ?", product.Id).Updates(model.Product{
			Category: dataProduct.Category,
			PricePerButir: dataProduct.PricePerButir,
			ExpiredDate: dataProduct.ExpiredDate,
			UpdatedAt: time.Now(),
		})

		if result.Error != nil {
			tx.Rollback()
			return errors.New("An error while update data product!"), http.StatusInternalServerError
		}

		if result.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Failed update data! Please try again later"), http.StatusBadRequest
		}

		// update stock
		resultStock := tx.Model(&model.Stock{}).Where("product_id = ?", product.Id).Updates(model.Stock{
			StockPerButir: dataProduct.Stock,
			LastUpdate: time.Now(),
		})

		if resultStock.Error != nil {
			tx.Rollback()
			return errors.New("An error while update data product stock!"), http.StatusInternalServerError
		}

		if result.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Failed update data! Please try again later"), http.StatusBadRequest
		}

		tx.Commit()
		return nil, http.StatusCreated
	}

	product = model.Product{
		Id: uuid.New().String(),
		Name: dataProduct.Name,
		Category: dataProduct.Category,
		PricePerButir: dataProduct.PricePerButir,
		ExpiredDate: dataProduct.ExpiredDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := tx.Create(&product).Error
	if err != nil {
		tx.Rollback()
		return errors.New("Failed create product! Please try again later!"), http.StatusInternalServerError
	}

	stock := model.Stock{
		Id: uuid.New().String(),
		StockPerButir: dataProduct.Stock,
		LastUpdate: time.Now(),
		ProductId: product.Id,
	}

	err = tx.Create(&stock).Error
	if err != nil {
		tx.Rollback()
		return errors.New("Failed create stock! Please try again later!"), http.StatusInternalServerError
	}

	tx.Commit()

	return nil, http.StatusOK
}

func (r *ProductRepository) GetProducts(limit int, page int, filter string)(model.ProductsResponseGet, error) {
	var products []model.Product	

	offset := (page - 1) * limit
	query := r.db.Model(&model.Product{})

	if filter != "" {
		query = query.Where("name LIKE ? OR category::text LIKE ?", "%"+ filter +"%", "%"+ strings.ToUpper(filter) +"%")
	}

	var totalRows int64
	query.Count(&totalRows)
	
	result := query.Limit(limit).Offset(offset).Find(&products)

	if result.Error != nil {
		return model.ProductsResponseGet{}, errors.New("An error while get data products. Please try again later!")
	}

	productResponse := make([]model.GetProducts, len(products)) 

	totalPages := int(math.Ceil(float64(totalRows) / float64(limit)))

	for i, product := range products {
		productResponse[i] = model.GetProducts{
			Id: product.Id,
			Name: product.Name,
			Category: product.Category,
			Price: product.PricePerButir,
			StatusExpired: helper.GetExpiredStatus(product.ExpiredDate),
			ExpiredDate: product.ExpiredDate,
		}
	}


	return model.ProductsResponseGet{
		TotalPages: int64(totalPages),
		Product: productResponse,
	}, nil
}