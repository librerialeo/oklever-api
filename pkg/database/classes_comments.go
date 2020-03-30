package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllClassesComments() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM classes_comments")
}
