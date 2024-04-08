package utils

import "golang.org/x/crypto/bcrypt"

func GetHashedPassword(password string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashBytes), err
}
