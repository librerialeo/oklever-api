package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllStudentsCourses queries for all studentsCourses
func (db *Database) GetAllStudentsCourses() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM students_courses")
}
