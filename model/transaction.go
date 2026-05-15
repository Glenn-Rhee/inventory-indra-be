package model

import "time"

type TransactionType string

const (
	TransactionIN	TransactionType = "IN"
	TransactionOUT 	TransactionType = "OUT"
)

type Transaction struct {
	Id              string 				`gorm:"primaryKey;type:varchar(255)"`
	ProductId       string 				`gorm:"type:varchar(255)"`
	TransactionType TransactionType 	`gorm:"not null"`
	Quantity        int    				`gorm:"not null"`
	CreatedAt       time.Time	

	Product			*Product			`gorm:"foreignKey:ProductId"`
}

type CreateTransaction struct {
	ProductId 		string				`json:"productId" binding:"required,min=5,max=100"`
	TransactionType TransactionType		`json:"transactionType" binding:"required,oneof=IN OUT"`
	Quantity		int					`json:"quantity" binding:"required,min=1"`
	Price			int					`json:"price" binding:"required,min=100"`
	ExpiredDate		*time.Time			`json:"expiredDate"`
}

type GetTransaction struct {
	TotalTransaction	int
	TotalRevenue		int
	TotalPurchase		int
	TotalPages			int
	Transactions		[]TransactionResponse
}

type TransactionResponse struct {
	Id 					string
	ProductName			string
	TransactionType		TransactionType
	Quantity			int
	Price				int64
	TotalPrice			int64
	TransactionDate		time.Time
}