package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string {
	return "postgres://postgres:123456@localhost:5432/school?sslmode=disable"
}

func NewConnection() (*sqlx.DB, error) {
	// Implementation for creating a new database connection
	dbSource := GetConnectionString()
	db, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		return nil, err
	}
	return db, nil
}
