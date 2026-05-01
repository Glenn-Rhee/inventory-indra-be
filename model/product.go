package model

import (
	"time"
)

type Product struct {
	Id				string 			`gorm:"primaryKey;type:varchar(255)"`
	Name 			string 			`gorm:"not null"`
	Category 		string 			`gorm:"not null"`
	PricePerButir	int 
	ExpiredDate 	time.Time		`gorm:"not null"`
	CreatedAt		time.Time
	UpdatedAt		time.Time

	Stock 			*Stock			`gorm:"foreignKey:ProductId;constraint:OnDelete:CASCADE"`
	Transaction		[]Transaction	`gorm:"foreignKey:ProductId"`
}