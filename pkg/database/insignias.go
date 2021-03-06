package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllInsignias queries for all insignias
func (db *Database) GetAllInsignias() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM insignias")
}
