package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	DB_HOST = "localhost"
	DB_PORT = "8881"
	DB_USER = "nbcamp-user"
	DB_PASS = "nbcamp-pass"
	DB_NAME = "nest"
)

func ConnectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
