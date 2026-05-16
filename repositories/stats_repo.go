package repositories

import (
	"errors"
	"inventory-indra/model"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type StatsRepository struct {
	db *gorm.DB
}

func NewStatsRepository(db *gorm.DB) *StatsRepository {
	return &StatsRepository{db: db}
}

func (r *StatsRepository) GetStatistik() (model.StastResponse, error, int) {
	var transaction []model.Transaction
	var totalTransaction int64
	query := r.db.Model(&model.Transaction{}).
		Joins("JOIN products ON products.id = transactions.product_id").
		Preload("Product")
	
	query.Count(&totalTransaction)
	
	result := query.Order("transactions.created_at ASC").Find(&transaction)
	if result.Error != nil {
		log.Println("Error while get data transaction:", result.Error.Error())
		return model.StastResponse{}, errors.New("An error while get data! Please try again later."), http.StatusInternalServerError
	}
	var totalRevenue int64 = 0
	var availableStock = 0

	chartMap := make(map[string]*model.DataChart)
	chartOrder := []string{}

	bestSellerMap := make(map[string]int)

	for _, tx := range transaction {
		dateKey := tx.CreatedAt.Format("2006-01-02")

		if _, isExist := chartMap[dateKey]; !isExist {
			chartMap[dateKey] = &model.DataChart{Date: dateKey}
			chartOrder = append(chartOrder, dateKey)
		}

		bestSellerMap[tx.Product.Name] += tx.Quantity

		amount := tx.Quantity * tx.Product.PricePerButir

		if tx.TransactionType == model.TransactionIN {
			chartMap[dateKey].In += amount
			totalRevenue += int64(amount)
		} else {
			chartMap[dateKey].Out += amount
		}
	}
	var bestSeller = ""
	maxQty := 0

	for productName, qty := range bestSellerMap {
		if qty > maxQty {
			maxQty = qty
			bestSeller = productName
		}
	}

	dataChart := make([]model.DataChart, 0, len(chartMap))
	for _, dateKey := range chartOrder {
		dataChart = append(dataChart, *chartMap[dateKey])
	}

	result = r.db.Model(&model.Stock{}).
		Select("SUM(stock_per_butir)").
		Scan(&availableStock)

	if result.Error != nil {
		log.Println("Error while get total stock:", result.Error.Error())
		return model.StastResponse{}, errors.New("An error while get data stock! Please try again later."), http.StatusInternalServerError
	}

	return model.StastResponse{
		TotalRevenue: totalRevenue,
		Stocks: int32(availableStock),
		DataChart: dataChart,
		BestSeller: bestSeller,
		TotalTransactions: int32(totalTransaction),
	}, nil, http.StatusOK
}