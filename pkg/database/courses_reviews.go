package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllCoursesReviews queries for all coursesReviews
func (db *Database) GetAllCoursesReviews() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM courses_reviews")
}
