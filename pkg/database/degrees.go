package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllDegrees queries for all degrees
func (db *Database) GetAllDegrees() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM degrees")
}
