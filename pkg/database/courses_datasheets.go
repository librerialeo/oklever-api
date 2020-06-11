package database

import (
	"github.com/jackc/pgx"
)

// GetAllCoursesDatasheets queries for all coursesDatasheets
func (db *Database) GetAllCoursesDatasheets() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM courses_datasheets")
}
