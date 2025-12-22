package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string {
	return "postgres://admin:secret@localhost:5432/mydb?sslmode=disable"
}

func NewConnection() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", GetConnectionString())
	if err != nil {
		return nil, err
	}

	fmt.Println("Database connected successfully")
	return db, nil
}

// func GetConnectionString() string {
// 	return "postgres://admin:secret@postgres:5432/mydb?sslmode=disable"
// }
// func NewConnection() (*sqlx.DB, error) {
// 	dbSource := GetConnectionString()
// 	dbCon, err := sqlx.Connect("postgres", dbSource)
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Println("Database is connected..")
// 	return dbCon, err
// }
