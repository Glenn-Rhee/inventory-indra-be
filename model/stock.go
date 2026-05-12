package model

import (
	"time"
)

type StatusStock string

const (
	StatusStockSafe 	StatusStock = "SAFE"
	StatusStockLow		StatusStock	= "LOW-STOCK"
	StatusStockSoldOut	StatusStock	= "SOLD-OUT"
)

type Stock struct {
	Id 				string 		`gorm:"primaryKey;type:varchar(255)"`
	StockPerButir	int 		`gorm:"not null"`
	LastUpdate		time.Time
	ProductId 		string		`gorm:"type:varchar(255)"`

	Product 		*Product	`gorm:"foreignKey:ProductId"`
}

type GetStocks struct {
	TotalProduct 		int64
	TotalLowStock		int
	TotalProductExpired	int
	Products			[]GetProductStock			 		
}

type GetProductStock struct {
	Id				string
	Name			string
	Stock			int
	StatusStock		StatusStock
	StatusExpired	ExpiredStatus
	ExpiredDate		time.Time
	LastUpdate		time.Time	
}