package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllClassesComments queries for all ClassesComments
func (db *Database) GetAllClassesComments() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM classes_comments")
}
