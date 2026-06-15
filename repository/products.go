package repository

import (
	"database/sql"
	"log"

	"github.com/VelVit24/projext/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) InsertProduct(instr *models.Product) error {
	err := r.db.QueryRow("insert into Products(name, category) values ($1, $2) returning id", instr.Name, instr.Id_category).Scan(&instr.Id)
	return err
}
func (r *ProductRepository) UpdateProduct(instr *models.Product) error {
	res, err := r.db.Exec("update Products set name=$1, category=$2 where id=$3", instr.Name, instr.Id_category, instr.Id)
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}
	return err
}
func (r *ProductRepository) DeleteProduct(id int) error {
	res, err := r.db.Exec("delete from Products where id=$1", id)
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}
	return err
}
func (r *ProductRepository) SelectProduct() ([]models.Product, error) {
	rows, err := r.db.Query("select id, name, category from Products")
	if err != nil {
		return nil, err
	}
	products := []models.Product{}
	for rows.Next() {
		product := models.Product{}
		err = rows.Scan(&product.Id, &product.Name, &product.Id_category)
		if err != nil {
			log.Print(err)
		}
		products = append(products, product)
	}
	return products, nil
}
