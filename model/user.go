package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Id        string `gorm:"primaryKey;type:varchar(255)"`
	Username  string `gorm:"not null"`
	Password  string `gorm:"not null"`
	ImgUrl    string
	CreatedAt time.Time

	Product     []Product     `gorm:"foreignKey:UserId"`
	Transaction []Transaction `gorm:"foreignKey:UserId"`
}

type Claims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	ImageUrl string `json:"imageUrl"`
	jwt.RegisteredClaims
}

type CreateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DataUser struct {
	Id       string
	Username string
	ImageUrl string
}

type UpdateUser struct {
	Id       string `json:"id"`
	ImageUrl string `json:"imageUrl"`
}
