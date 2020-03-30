package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllCoursesDatasheets queries for all coursesDatasheets
func (db *Database) GetAllCoursesDatasheets() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM courses_datasheets")
}
