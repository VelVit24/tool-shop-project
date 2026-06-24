package repository

import (
	"database/sql"
	"log"

	"github.com/VelVit24/projext/dto"
	"github.com/VelVit24/projext/models"
)

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) InsertCart(id_user int, cart *models.Cart) error {
	_, err := r.db.Exec(`
	insert into cart_items(id_user, id_product, amount) 
	values ($1, $2, $3)`,
		id_user, cart.Id_product, cart.Amount)
	return err
}

func (r *CartRepository) UpdateCart(id_user int, id int, cart *models.Cart) error {
	res, err := r.db.Exec("update cart_items set amount=$1 where id_user=$2 and id_product=$3", cart.Amount, id_user, id)
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}
	return err
}
func (r *CartRepository) DeleteCart(id_user, id int) error {
	res, err := r.db.Exec("delete from cart_items where id_user=$1 and id_product=$2", id_user, id)
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}
	return err
}

func (r *CartRepository) SelectCart(id_user int) ([]dto.CartItems, error) {
	rows, err := r.db.Query("select id_product, name, price, stock, image_count, amount from cart_items c left outer join products p on c.id_product = p.id where id_user=$1", id_user)
	if err != nil {
		return nil, err
	}
	items := []dto.CartItems{}
	for rows.Next() {
		item := dto.CartItems{}
		err := rows.Scan(&item.Id_product, &item.Name, &item.Price, &item.Stock, &item.Image_count, &item.Amount)
		if err != nil {
			log.Println(err)
		}
		if item.Amount > item.Stock {
			item.IsInStock = false
		} else {
			item.IsInStock = true
		}
		items = append(items, item)
	}
	return items, err
}

func (r *CartRepository) SelectProductStock(id_prod int) (int, error) {
	var stock int
	err := r.db.QueryRow("select stock from products where id=$1", id_prod).Scan(&stock)
	return stock, err
}

func (r *CartRepository) DeleteAllCart(id_user int) error {
	_, err := r.db.Exec("delete from cart_items where id_user=$1", id_user)
	return err
}
