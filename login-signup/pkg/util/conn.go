package util

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func GetDBConnection() (*sql.DB, error) {
	connStr := "user=postgres password=admin host=localhost port=7700 dbname=catea_main sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}
