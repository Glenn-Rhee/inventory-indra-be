package model

import (
	"time"
)

type Stock struct {
	Id 				string 		`gorm:"primaryKey;type:varchar(255)"`
	StockPerButir	int 		`gorm:"not null"`
	LastUpdate		time.Time
	ProductId 		string		`gorm:"type:varchar(255)"`

	Product 		*Product	`gorm:"foreignKey:ProductId"`
}