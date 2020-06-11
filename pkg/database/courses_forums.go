package database

import (
	"github.com/jackc/pgx"
)

// GetAllCoursesForums queries for all coursesForums
func (db *Database) GetAllCoursesForums() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM courses_forums")
}
