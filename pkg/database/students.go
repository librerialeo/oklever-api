package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllStudents queries for all students
func (db *Database) GetAllStudents() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM students")
}
