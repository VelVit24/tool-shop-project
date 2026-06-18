package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/VelVit24/projext/dto"
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

func (r *ProductRepository) SelectProducts(filter dto.ProductFiler) ([]models.Product, error) {
	query := "select id, name, description, price, stock, image_url, id_category from products where 1=1"
	args := []any{}
	ind := 1
	if filter.CategoryID != nil {
		query += fmt.Sprintf(" and id_category=$%d", ind)
		args = append(args, *filter.CategoryID)
		ind++
	}
	if filter.PriceFrom != nil {
		query += fmt.Sprintf(" and price >= $%d", ind)
		args = append(args, *filter.PriceFrom)
		ind++
	}
	if filter.PriceTo != nil {
		query += fmt.Sprintf(" and price <= $%d", ind)
		args = append(args, *filter.PriceTo)
		ind++
	}
	if filter.InStock != nil {
		query += " and stock >= 0"
	}

	if filter.Search != nil {
		query += fmt.Sprintf(" and name ilike $%d or description ilike $%d", ind, ind+1)
		search := "%" + *filter.Search + "%"
		args = append(args, search, search)
		ind += 2
	}
	switch filter.Sort {
	case "price_asc":
		query += " order by price asc"
	case "price_desc":
		query += " order by price desc"
	case "name_asc":
		query += " order by name asc"
	case "name_desc":
		query += " order by name desc"
	}
	offset := (filter.Page - 1) * filter.Limit
	query += fmt.Sprintf(" offset $%d limit $%d", ind, ind+1)
	args = append(args, offset, filter.Limit)

	rows, err := r.db.Query(query, args...)
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
