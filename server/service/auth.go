package service

import (
	"errors"
	"net/mail"
	"strings"

	"github.com/VelVit24/projext/models"
	"github.com/VelVit24/projext/repository"
	"github.com/golang-jwt/jwt/v5"
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

func (s *AuthService) CreateUser(user *models.User) error {
	if !ValidateEmail(user.Email) {
		return errors.New("invalid email")
	}
	if len(user.Password) < 8 {
		return errors.New("invalid password length")
	}
	normalizedPhone, err := NormalizePhone(user.Phone)
	if err != nil {
		return err
	}
	user.Phone = normalizedPhone

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

func (s *AuthService) CheckEmail(email string) error {
	if !ValidateEmail(email) {
		return errors.New("invalid email")
	}
	isUnique := s.repo.CheckEmailUnique(email)
	if !isUnique {
		return errors.New("email already exists")
	}
	return nil
}

func (s *AuthService) CheckPhoneUnique(phone string) error {
	normalizedPhone, err := NormalizePhone(phone)
	if err != nil {
		return err
	}
	isUnique := s.repo.CheckPhoneUnique(normalizedPhone)
	if !isUnique {
		return errors.New("phone already exists")
	}
	return nil
}

func ValidateEmail(email string) bool {

	addr, err := mail.ParseAddress(email)

	if err != nil {
		return false
	}

	parts := strings.Split(addr.Address, "@")

	if len(parts) != 2 {
		return false
	}

	domain := parts[1]

	return strings.Contains(domain, ".")
}

func NormalizePhone(phone string) (string, error) {
	normalized := ""
	for _, ch := range phone {
		if ch >= '0' && ch <= '9' {
			normalized += string(ch)
		}
	}
	if len(normalized) != 11 {
		return "", errors.New("invalid phone")
	}
	return normalized, nil
}
