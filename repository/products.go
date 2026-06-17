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

func (r *ProductRepository) InsertProduct(prod *models.Product) error {
	err := r.db.QueryRow("insert into products(name, description, price, stock, image_url, id_category) values ($1, $2, $3, $4, $5, $6) returning id",
		prod.Name, prod.Description, prod.Price, prod.Stock, prod.Image_url, prod.Id_category).Scan(&prod.Id)
	return err
}
func (r *ProductRepository) UpdateProduct(prod *models.Product) error {
	res, err := r.db.Exec("update products set name=$1, description=$2, price=$3, stock=$4, image_url=$5, id_category=$6 where id=$7",
		prod.Name, prod.Description, prod.Price, prod.Stock, prod.Image_url, prod.Id_category, prod.Id)
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}
	return err
}
func (r *ProductRepository) DeleteProduct(id int) error {
	res, err := r.db.Exec("delete from products where id=$1", id)
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}
	return err
}

func (r *ProductRepository) SelectProducts(page, limit int) ([]models.Product, error) {
	rows, err := r.db.Query("select id, name, description, price, stock, image_url, id_category from products offset $1 limit $2", (page-1)*limit, limit)
	if err != nil {
		return nil, err
	}
	products := []models.Product{}
	for rows.Next() {
		product := models.Product{}
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Stock, &product.Image_url, &product.Id_category)
		if err != nil {
			log.Print(err)
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *ProductRepository) SelectProductId(id int) (models.Product, error) {
	product := models.Product{}
	err := r.db.QueryRow("select id, name, description, price, stock, image_url, id_category from products where id=$1", id).
		Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Stock, &product.Image_url, &product.Id_category)
	return product, err
}
