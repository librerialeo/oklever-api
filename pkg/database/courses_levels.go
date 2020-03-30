package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllCoursesLevels queries for all coursesLevels
func (db *Database) GetAllCoursesLevels() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM courses_levels")
}
