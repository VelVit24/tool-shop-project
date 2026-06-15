package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Id int `json:"user_id"`
	jwt.RegisteredClaims
}

func HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	return string(bytes), err
}
func CheckPassword(pass, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err
}

func GenToken(id int) (string, error) {
	key := []byte("key")
	claims := Claims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
