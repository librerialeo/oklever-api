package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllSynchronousClasses() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM synchronous_classes")
}
