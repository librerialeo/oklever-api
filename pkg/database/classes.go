package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllClasses queries for all classes
func (db *Database) GetAllClasses() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM classes")
}
