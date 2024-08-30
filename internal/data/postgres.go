package data

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func getConnection() (*sql.DB, error) {
	uri := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	return sql.Open("postgres", uri)
}

func MakeMigration(db *sql.DB) error {
	b, err := os.ReadFile("./database/models.sql")
	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))
	if err != nil {
		return err
	}

	return rows.Close()
}
