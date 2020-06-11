package database

import (
	"github.com/jackc/pgx"
)

// GetAllCoursesLevels queries for all coursesLevels
func (db *Database) GetAllCoursesLevels() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM courses_levels")
}
