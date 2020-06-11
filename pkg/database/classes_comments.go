package database

import (
	"github.com/jackc/pgx"
)

// GetAllClassesComments queries for all ClassesComments
func (db *Database) GetAllClassesComments() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM classes_comments")
}
