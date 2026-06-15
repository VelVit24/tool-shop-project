package repository

import (
	"database/sql"
	"log"

	"github.com/VelVit24/projext/models"
)

func InsertInstrument(db *sql.DB, instr *models.Instrument) error {
	err := db.QueryRow("insert into Instruments(name, category) values ($1, $2) returning id", instr.Name, instr.Id_category).Scan(&instr.Id)
	return err
}
func UpdateInstrument(db *sql.DB, instr *models.Instrument) error {
	res, err := db.Exec("update Instruments set name=$1, category=$2 where id=$3", instr.Name, instr.Id_category, instr.Id)
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}
	return err
}
func DeleteInstrument(db *sql.DB, id int) error {
	res, err := db.Exec("delete from Instruments where id=$1", id)
	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}
	return err
}
func SelectInstrument(db *sql.DB) ([]models.Instrument, error) {
	rows, err := db.Query("select id, name, category from Instruments")
	if err != nil {
		return nil, err
	}
	instrs := []models.Instrument{}
	for rows.Next() {
		instr := models.Instrument{}
		err = rows.Scan(&instr.Id, &instr.Name, &instr.Id_category)
		if err != nil {
			log.Print(err)
		}
		instrs = append(instrs, instr)
	}
	return instrs, nil
}
