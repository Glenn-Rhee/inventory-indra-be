package model

import "time"

type User struct {
	Id        string 		`gorm:"primaryKey;type:varchar(255)"`
	Username  string 		`gorm:"not null"`
	Password  string 		`gorm:"not null"`
	ImageUrl  string
	CreatedAt time.Time
}