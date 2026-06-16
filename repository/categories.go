package repository

import (
	"database/sql"
	"log"

	"github.com/VelVit24/projext/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) InsertCategory(cat *models.Category) error {
	err := r.db.QueryRow("insert into categories(name) values ($1) returning id", cat.Name).Scan(&cat.Id)
	return err
}
func (r *CategoryRepository) UpdateCategory(cat *models.Category) error {
	res, err := r.db.Exec("update categories set name=$1 where id=$2", cat.Name, cat.Id)
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}
	return err
}
func (r *CategoryRepository) DeleteCategory(id int) error {
	res, err := r.db.Exec("delete from categories where id=$1", id)
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}
	return err
}
func (r *CategoryRepository) SelectCategories() ([]models.Category, error) {
	rows, err := r.db.Query("select id, name from categories")
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
