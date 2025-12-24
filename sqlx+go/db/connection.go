package db

import (
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionStr() string {
	return "postgres://admin:secret@localhost:5432/article?sslmode=disable"
}
func NewConnection() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", GetConnectionStr())
	if err != nil {
		return nil, errors.New("Db is not connected!")
	}
	return db, nil
}
