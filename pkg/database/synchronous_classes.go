package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllSynchronousClasses queries for all synchronousClasses
func (db *Database) GetAllSynchronousClasses() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM synchronous_classes")
}
