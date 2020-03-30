package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllModules queries for all modules
func (db *Database) GetAllModules() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM modules")
}
