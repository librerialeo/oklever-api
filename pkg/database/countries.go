package database

import (
	"github.com/jackc/pgx"
)

// GetAllCountries queries for all countries
func (db *Database) GetAllCountries() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM countries")
}
