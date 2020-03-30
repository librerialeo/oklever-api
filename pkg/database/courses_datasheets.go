package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllCoursesDatasheets() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM courses_datasheets")
}
