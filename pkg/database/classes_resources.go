package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllClassesResources() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM classes_resources")
}
