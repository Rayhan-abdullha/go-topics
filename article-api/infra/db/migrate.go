package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDB(db *sqlx.DB, dir string) error {
	migrations := migrate.FileMigrationSource{
		Dir: dir,
	}

	n, err := migrate.Exec(
		db.DB,      // *sql.DB
		"postgres", // dialect
		migrations,
		migrate.Up, // or migrate.Down
	)
	if err != nil {
		return err
	}

	log.Printf("Applied %d migrations\n", n)
	return nil
}
