package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllClassesResources queries for all ClassesResources
func (db *Database) GetAllClassesResources() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM classes_resources")
}
