package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllTestClasses queries for all testClasses
func (db *Database) GetAllTestClasses() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM test_classes")
}
