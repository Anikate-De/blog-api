package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

// TODO: Load key from env
const key = "some secret key"

func GenerateJWT(email string, id int64) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"id":    id,
		},
	)

	return token.SignedString([]byte(key))
}
