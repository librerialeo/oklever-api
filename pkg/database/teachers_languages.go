package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllTeachersLanguages() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM teachers_languages")
}
