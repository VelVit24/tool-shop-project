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
	err := r.db.QueryRow(`
	insert into 
	users(email, password, phone, first_name, last_name) 
	values ($1, $2, $3, $4, $5) 
	returning id, role`,
		user.Email, user.Password,
		user.Phone, user.FirstName,
		user.LastName).Scan(&user.Id, &user.Role)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthRepository) CheckUser(user *models.User) (string, error) {
	var hash string
	err := r.db.QueryRow(`
	select id, password, role 
	from users 
	where email=$1 or phone=$2`,
		user.Email, user.Phone).Scan(&user.Id, &hash, &user.Role)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (r *AuthRepository) CheckEmailUnique(email string) bool {
	id := -1
	err := r.db.QueryRow("select id from users where email=$1",
		email).Scan(&id)
	if err != nil || id == -1 {
		return true
	}
	return false
}
func (r *AuthRepository) CheckPhoneUnique(phone string) bool {
	id := -1
	err := r.db.QueryRow("select id from users where phone=$1",
		phone).Scan(&id)
	if err != nil || id == -1 {
		return true
	}
	return false
}

func (r *AuthRepository) UpdateUser(user *models.User) error {
	_, err := r.db.Exec(`
	update users set
	email=$1, password=$2, phone=$3, first_name=$4, last_name=$5
	where id=$6`,
		user.Email, user.Password,
		user.Phone, user.FirstName,
		user.LastName, user.Id)
	if err != nil {
		return err
	}
	return nil
}
