package database

import (
	"github.com/jackc/pgx"
)

// GetAllCoursesForumsComments queries for all coursesForumsComments
func (db *Database) GetAllCoursesForumsComments() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM courses_forums_comments")
}
