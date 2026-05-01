package repositories

import (
	"inventory-indra/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return  &UserRepository{db: db}
}

func (r *UserRepository) FindUser() (model.User, error) {
	var user model.User
	err := r.db.First(&user).Error
	return user, err
}