package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectToDB() (*sql.DB, error) {
	connStr := "host=localhost user=postgres password=toor dbname=sr_test port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the database")
	return db, nil
}
