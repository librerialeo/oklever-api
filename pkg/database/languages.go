package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllLanguages queries for all languages
func (db *Database) GetAllLanguages() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM languages")
}
