package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Id        string 		`gorm:"primaryKey;type:varchar(255)"`
	Username  string 		`gorm:"not null"`
	Password  string 		`gorm:"not null"`
	ImgUrl  string
	CreatedAt time.Time
}

type Claims struct {
	Id 			string `json:"id"`
	Username	string `json:"username"`
	jwt.RegisteredClaims
}

type CreateUser struct {
	Username 	string 	`json:"username"`
	Password 	string	`json:"password"`
}

type DataUser struct {
	Id			string
	Username 	string
	ImageUrl 	string
}