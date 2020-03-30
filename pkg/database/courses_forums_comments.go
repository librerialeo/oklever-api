package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllCoursesForumsComments queries for all coursesForumsComments
func (db *Database) GetAllCoursesForumsComments() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM courses_forums_comments")
}
