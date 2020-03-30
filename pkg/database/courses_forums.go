package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllCoursesForums queries for all coursesForums
func (db *Database) GetAllCoursesForums() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM courses_forums")
}
