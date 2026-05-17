package model

import "time"

type StastResponse struct {
	TotalRevenue      int64
	TotalTransactions int32
	Stocks            int32
	BestSeller        string
	DataChart         []DataChart
}

type DataChart struct {
	Date string // YYYY-MM-DD
	In   int
	Out  int
}

type DataMedicineResponse struct {
	Id            string
	ProductName   string
	Category      Category
	PricePerButir int
	ExpiredDate   time.Time
	StockPerButir int
	LastUpdate    time.Time
}