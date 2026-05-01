package lib

import "golang.org/x/crypto/bcrypt"

func GenereteHash(password string) (hashedPass string, err error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}