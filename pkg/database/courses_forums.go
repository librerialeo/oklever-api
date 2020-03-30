package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllCoursesForums() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM courses_forums")
}
