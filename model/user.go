package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Id        string 		`gorm:"primaryKey;type:varchar(255)"`
	Username  string 		`gorm:"not null"`
	Password  string 		`gorm:"not null"`
	ImageUrl  string
	CreatedAt time.Time
}

type Claims struct {
	Id 			string `json:"id"`
	Username	string `json:"username"`
	jwt.RegisteredClaims
}