package repository

import (
	"database/sql"
	"log"

	"github.com/VelVit24/projext/models"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) InsertOrder(db *sql.DB, id_user int, order models.Order) error {
	id := 0
	err := db.QueryRow("insert into Orders(id_user, description) values ($1, $2) returning id", id_user).Scan(&id)
	if err != nil {
		return err
	}
	carts, err := SelectCart(db, id_user)
	for _, cart := range carts {
		_, err := db.Exec("insert into OrderItems(id_order, id_instr, amount) values ($1, $2, $3)", id, cart.Id_product, cart.Amount)
		if err != nil {
			log.Println(err)
		}
	}
	_, err = db.Exec("delete from CartItems where id_user=$1", id_user)
	return err
}

//	func UpdateOrder(db *sql.DB, id_user int, order models.Order) error {
//		res, err := db.Exec("update Orders set description=$1 where id_user=$2 and id_instr=$3", cart.Amount, id_user, cart.Id_instr)
//		if rows, _ := res.RowsAffected(); rows == 0 {
//			return sql.ErrNoRows
//		}
//		return err
//	}
func DeleteCart(db *sql.DB, id_user, id int) error {
	res, err := db.Exec("delete from CartItems where id_user=$1 and id_instr=$2", id_user, id)
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}
	return err
}

func SelectCart(db *sql.DB, id_user int) ([]models.Cart, error) {
	rows, err := db.Query("select id_instr, amount from CartItems where id_user=$1", id_user)
	if err != nil {
		return nil, err
	}
	carts := []models.Cart{}
	for rows.Next() {
		cart := models.Cart{}
		err := rows.Scan(&cart.Id_product, &cart.Amount)
		if err != nil {
			log.Println(err)
		}
		carts = append(carts, cart)
	}
	return carts, err
}
