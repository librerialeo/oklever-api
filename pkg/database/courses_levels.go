package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllCoursesLevels() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM courses_levels")
}
