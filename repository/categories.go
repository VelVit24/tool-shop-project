package repository

import (
	"database/sql"
	"log"

	"github.com/VelVit24/projext/models"
)

func InsertCategory(db *sql.DB, cat *models.Category) error {
	err := db.QueryRow("insert into Categories(name) values ($1) returning id", cat.Name).Scan(&cat.Id)
	return err
}
func UpdateCategory(db *sql.DB, cat *models.Category) error {
	res, err := db.Exec("update Categories set name=$1 where id=$2", cat.Name, cat.Id)
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}
	return err
}
func DeleteCategory(db *sql.DB, id int) error {
	res, err := db.Exec("delete from Categories where id=$1", id)
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}
	return err
}
func SelectCategories(db *sql.DB) ([]models.Category, error) {
	rows, err := db.Query("select id, name from Categories")
	if err != nil {
		return nil, err
	}
	cats := []models.Category{}
	for rows.Next() {
		cat := models.Category{}
		err = rows.Scan(&cat.Id, &cat.Name)
		if err != nil {
			log.Print(err)
		}
		cats = append(cats, cat)
	}
	return cats, nil
}
