package repository

import (
	"database/sql"

	"github.com/VelVit24/projext/models"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) InsertUser(user *models.User) error {
	err := r.db.QueryRow("insert into users(email, password) values ($1, $2) returning id, role", user.Email, user.Password).Scan(&user.Id, &user.Role)
	if err != nil {
		return err
	}
	return err
}

func (r *AuthRepository) CheckUser(user *models.User) (string, error) {
	var hash string
	err := r.db.QueryRow("select id, password, role from users where email=$1", user.Email).Scan(&user.Id, &hash, &user.Role)
	if err != nil {
		return "", err
	}
	return hash, nil
}
