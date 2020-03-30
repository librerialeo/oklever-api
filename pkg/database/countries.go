package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllCountries queries for all countries
func (db *Database) GetAllCountries() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM countries")
}
