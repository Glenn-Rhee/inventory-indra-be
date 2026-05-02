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
	return &UserRepository{db: db}
}

func (r *UserRepository) LoginUser(dataUser model.CreateUser) (model.DataUser, error) {
	var user model.User
	result := r.db.Where("username = ?", dataUser.Username).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound){
		return model.DataUser{}, errors.New("Username is not found!")
	}

	err := lib.ComparePass(dataUser.Password, user.Password)
	if err != nil {
		return model.DataUser{}, errors.New("Password doesn't match!")
	}

	return model.DataUser{
		Id: user.Id,
		Username: user.Username,
		ImageUrl: user.ImgUrl,
	}, nil
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

func(r *UserRepository) GetOneUser(userId string) (model.DataUser, error){
	var user model.User

	result := r.db.Where("id = ?", userId).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound){
		return model.DataUser{}, errors.New("User is not found!")
	}

	return model.DataUser{
		Id: user.Id,
		Username: user.Username,
		ImageUrl: user.ImgUrl,
	}, nil
}