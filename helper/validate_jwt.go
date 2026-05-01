package helper

import (
	"errors"
	"inventory-indra/model"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func ValidateJWT(tokenString string) (*model.Claims, error) {
	if err:= godotenv.Load(); err !=nil {
		log.Println("Error load dotenv file:", err.Error())
	}

	secret := os.Getenv("SECRET_KEY")

	token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return  nil, err
	}

	if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Invalid Token")
}