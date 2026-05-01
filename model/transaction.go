package model

import "time"

type Transaction struct {
	Id              string 		`gorm:"primaryKey;type:varchar(255)"`
	ProductId       string 		`gorm:"type:varchar(255)"`
	TransactionType string 		`gorm:"not null"`
	Quantity        int    		`gorm:"not null"`
	CreatedAt       time.Time	

	Product			*Product	`gorm:"foreignKey:ProductId"`
}