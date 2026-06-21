package service

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	return string(bytes), err
}
func CheckPassword(pass, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err
}

func GenToken(id int, role string) (string, error) {
	log.Println(id, role)
	key := []byte("key")
	claims := Claims{
		Id:   id,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
