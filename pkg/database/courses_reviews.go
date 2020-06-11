package database

import (
	"github.com/jackc/pgx"
)

// GetAllCoursesReviews queries for all coursesReviews
func (db *Database) GetAllCoursesReviews() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM courses_reviews")
}
