package lib

import "golang.org/x/crypto/bcrypt"

func GenereteHash(password string) (hashedPass string, err error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

func ComparePass(password string, hashPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
}