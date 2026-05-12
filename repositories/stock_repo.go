package repositories

import (
	"errors"
	"inventory-indra/helper"
	"inventory-indra/model"
	"net/http"

	"gorm.io/gorm"
)

type StockRepository struct {
	db *gorm.DB
}

type GetStocksParams struct {
	Limit 	int
	Page	int
	Filter	string
}

func NewStockRepository(db *gorm.DB) *StockRepository {
	return &StockRepository{db: db}
}

func (r *StockRepository) GetStocks(params GetStocksParams) (model.GetStocks, error, int) {
	var products []model.Product
	var totalLowStock = 0
	var totalProductExpired = 0
	offset := (params.Page - 1) * params.Limit
	query := r.db.Model(&model.Product{}).
		Joins("Stock").
		Where("products.is_active = ?", true)
	
	if params.Filter != "" {
		query = query.Where("products.name ILIKE ?", "%" + params.Filter + "%")
	}

	var totalStock int64
	query.Count(&totalStock)
	result := query.Order("products.name ASC").Limit(params.Limit).Offset(offset).Find(&products).Count(&totalStock)

	if result.Error != nil {
		return model.GetStocks{}, errors.New("An error while get data"), http.StatusInternalServerError
	}

	productsResponse := make([]model.GetProductStock, len(products))
	
	for i, product := range products {
		statusExpired := helper.GetExpiredStatus(product.ExpiredDate)
		productsResponse[i] = model.GetProductStock{
			Id: product.Stock.Id,
			Name: product.Name,
			Stock: product.Stock.StockPerButir,
			StatusExpired: statusExpired,
			ExpiredDate: product.ExpiredDate,
			StatusStock: helper.GetStatusStock(product.Stock.StockPerButir),
			LastUpdate: product.Stock.LastUpdate,
		}

		if statusExpired == model.StatusExpired {
			totalProductExpired++
		}

		if product.Stock.StockPerButir < 10 {
			totalLowStock++
		}
	}

	return model.GetStocks{
		TotalProduct: totalStock,
		TotalLowStock: totalLowStock,
		TotalProductExpired: totalProductExpired,
		Products: productsResponse,
	}, nil, http.StatusOK
}