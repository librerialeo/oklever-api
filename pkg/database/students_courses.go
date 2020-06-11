package database

import (
	"github.com/jackc/pgx"
)

// GetAllStudentsCourses queries for all studentsCourses
func (db *Database) GetAllStudentsCourses() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM students_courses")
}
