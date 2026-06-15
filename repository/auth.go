package repository

import (
	"database/sql"

	"github.com/VelVit24/projext/models"
	"github.com/VelVit24/projext/service"
)

func InsertUser(db *sql.DB, user models.User) (int, error) {
	var id int
	hash, err := service.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	err = db.QueryRow("insert into Users(email, password) values ($1, $2) returning id", user.Email, hash).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, sql.ErrNoRows
		}
		return 0, err
	}
	return id, err
}

func CheckUser(db *sql.DB, user *models.User) error {
	var hash string
	err := db.QueryRow("select id, password from Users where email=$1", user.Email).Scan(&user.Id, &hash)
	if err != nil {
		return err
	}
	return service.CheckPassword(user.Password, hash)
}
