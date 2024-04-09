package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	Email string
	Id    int64
	jwt.RegisteredClaims
}

// TODO: Load key from env
const key = "some secret key"

func GenerateJWT(email string, id int64) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		UserClaims{
			Email: email,
			Id:    id,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
			},
		},
	)

	return token.SignedString([]byte(key))
}

func ParseJWT(tokenString string) (int64, error) {
	var claims UserClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return 0, err
	}

	parsedClaims, ok := token.Claims.(*UserClaims)
	if !ok {
		return 0, errors.New("unknown claims type, cannot proceed")
	}

	if parsedClaims.ExpiresAt.Time.Before(time.Now()) {
		return 0, errors.New("token expired")
	}

	return parsedClaims.Id, nil

}
