package repositories

import (
	"errors"
	"inventory-indra/lib"
	"inventory-indra/model"
	"time"

	"github.com/google/uuid"
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

func (r *UserRepository) CreateUser(dataUser model.CreateUser) (error) {
	var user model.User

	var existingUser model.User
	result := r.db.Where("username = ?", dataUser.Username).First(&existingUser)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("Username already registered!")
	}
	
	passHashed, err := lib.GenereteHash(dataUser.Password)
	if err != nil {
		return errors.New("An error while create user")
	}

	user = model.User{
		Id: uuid.New().String(),
		Username: dataUser.Username,
		Password: passHashed,
		ImgUrl: "",
		CreatedAt: time.Now(),
	}

	err = r.db.Create(user).Error
	if err != nil {
		return errors.New("An error while creating user!")
	}
	
	return nil
}