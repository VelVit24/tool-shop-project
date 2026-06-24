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
	err := r.db.QueryRow("insert into products(name, description, price, stock, image_count, id_category, slug) values ($1, $2, $3, $4, $5, $6, $7) returning id",
		prod.Name, prod.Description, prod.Price, prod.Stock, prod.Image_count, prod.Id_category, prod.Slug).Scan(&prod.Id)
	return err
}
func (r *ProductRepository) UpdateProduct(prod *models.Product) error {
	res, err := r.db.Exec("update products set name=$1, description=$2, price=$3, stock=$4, image_count=$5, id_category=$6 where id=$7",
		prod.Name, prod.Description, prod.Price, prod.Stock, prod.Image_count, prod.Id_category, prod.Id)
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

func (r *ProductRepository) SelectProducts(filter dto.ProductFiler) (dto.ProductsResponce, error) {
	query := " from products where 1=1"
	args := []any{}
	ind := 1
	if filter.CategorySlug != nil {
		category_id := 0
		r.db.QueryRow("select id from categories where slug=$1", filter.CategorySlug).Scan(&category_id)
		query += fmt.Sprintf(" and id_category=$%d", ind)
		args = append(args, category_id)
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

	var total int
	countQuery := "select count(*)" + query
	err := r.db.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return dto.ProductsResponce{}, err
	}

	selectQuery := "select id, name, description, price, stock, image_count, id_category, slug" + query
	switch filter.Sort {
	case "price_asc":
		selectQuery += " order by price asc"
	case "price_desc":
		selectQuery += " order by price desc"
	case "name_asc":
		selectQuery += " order by name asc"
	case "name_desc":
		selectQuery += " order by name desc"
	}

	offset := (filter.Page - 1) * filter.Limit
	selectQuery += fmt.Sprintf(" offset $%d limit $%d", ind, ind+1)
	args = append(args, offset, filter.Limit)

	rows, err := r.db.Query(selectQuery, args...)
	if err != nil {
		return dto.ProductsResponce{}, err
	}
	products := []models.Product{}
	for rows.Next() {
		product := models.Product{}
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Stock, &product.Image_count, &product.Id_category, &product.Slug)
		if err != nil {
			log.Print(err)
		}
		products = append(products, product)
	}
	responce := dto.ProductsResponce{
		Page:     filter.Page,
		Limit:    filter.Limit,
		Products: products,
		Total:    total,
	}
	return responce, nil
}

func (r *ProductRepository) SelectProductSlug(slug string) (models.Product, error) {
	product := models.Product{}
	err := r.db.QueryRow("select id, name, description, price, stock, image_count, id_category, slug from products where slug=$1", slug).
		Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Stock, &product.Image_count, &product.Id_category, &product.Slug)
	return product, err
}

func (r *ProductRepository) AddProductImageCount(slug string, count int) error {
	res, err := r.db.Exec("update products set image_count=image_count+$1 where slug=$2", count, slug)
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}
	return err
}

func (r *ProductRepository) CheckSlug(slug string) error {
	id := -1
	err := r.db.QueryRow("select id from product where slug=$1", slug).Scan(&id)
	if err != nil {
		return err
	}
	if id != -1 {
		return fmt.Errorf("slug already exists")
	}
	return nil
}
