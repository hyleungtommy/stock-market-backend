package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func GetPostgres() (*sql.DB, error) {
	connectionString := "user=postgres password=root dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	return db, err
}
