package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllCoursesReviews() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM courses_reviews")
}
