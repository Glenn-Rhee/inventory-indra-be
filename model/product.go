package model

import (
	"time"
)


type ExpiredStatus string

const (
	StatusSafe    ExpiredStatus = "SAFE"
	StatusWarning ExpiredStatus = "WARNING"
	StatusExpired ExpiredStatus = "EXPIRED"
)

type Category string

const (
	CategoryMedicine Category = "MEDICINE"
	CategoryEssentials Category = "ESSENTIALS" 
)

type Product struct {
	Id				string 			`gorm:"primaryKey;type:varchar(255)"`
	Name 			string 			`gorm:"not null"`
	Category 		Category 		`gorm:"not null"`
	PricePerButir	int 
	ExpiredDate 	time.Time		`gorm:"not null"`
	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		time.Time
	IsActive		bool

	Stock 			*Stock			`gorm:"foreignKey:ProductId;constraint:OnDelete:CASCADE"`
	Transaction		[]Transaction	`gorm:"foreignKey:ProductId"`
}

type CreateProduct struct {
	Name 			string		`json:"name" binding:"required,min=5,max=100"`
	Category		Category	`json:"category" binding:"required,oneof=MEDICINE ESSENTIALS"`
	Stock			int			`json:"stock" binding:"required,min=1"`
	PricePerButir	int			`json:"pricePerButir" binding:"required,min=1000"`
	ExpiredDate		time.Time	`json:"expiredDate" binding:"required"`
}

type ProductsResponseGet struct {
	TotalPages 	int64
	Product	 	[]GetProducts
}

type GetProducts struct {
	Id 				string
	Name 			string
	Category 		Category
	Price			int
	StatusExpired	ExpiredStatus
	ExpiredDate		time.Time
}