package database

import (
	"github.com/jackc/pgx"
)

// GetAllStudents queries for all students
func (db *Database) GetAllStudents() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM students")
}
