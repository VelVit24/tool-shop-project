package service

import (
	"log"
	"time"

	"github.com/VelVit24/projext/models"
	"github.com/VelVit24/projext/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

type Claims struct {
	Id   int    `json:"user_id"`
	Role string `json:"role"`
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

func (s *AuthService) CreateUser(user *models.User) error {
	hash, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	err = s.repo.InsertUser(user)
	return err
}

func (s *AuthService) CheckUser(user *models.User) error {
	hash, err := s.repo.CheckUser(user)
	if err != nil {
		return err
	}
	err = CheckPassword(user.Password, hash)
	return err
}
