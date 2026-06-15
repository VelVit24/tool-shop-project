package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnDB() *sql.DB {
	connStr := "user=postgres password=080907 dbname=projectdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
