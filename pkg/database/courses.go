package database

import (
	"github.com/jackc/pgx"
)

// GetAllCourses queries for all courses
func (db *Database) GetAllCourses() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM courses")
}
