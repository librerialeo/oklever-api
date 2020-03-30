package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllSynchronousClassesResources queries for all synchronousClassesResources
func (db *Database) GetAllSynchronousClassesResources() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM synchronous_classes_resources")
}
