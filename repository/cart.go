package repository

import (
	"database/sql"
	"log"

	"github.com/VelVit24/projext/models"
)

func InsertCart(db *sql.DB, id_user int, cart models.Cart) error {
	_, err := db.Exec("insert into CartItems(id_user, id_instr, amount) values ($1, $2, $3)", id_user, cart.Id_instr, cart.Amount)
	return err
}

func UpdateCart(db *sql.DB, id_user int, cart models.Cart) error {
	res, err := db.Exec("update CartItems set amount=$1 where id_user=$2 and id_instr=$3", cart.Amount, id_user, cart.Id_instr)
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}
	return err
}
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
		err := rows.Scan(&cart.Id_instr, &cart.Amount)
		if err != nil {
			log.Println(err)
		}
		carts = append(carts, cart)
	}
	return carts, err
}
