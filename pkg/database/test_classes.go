package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllTestClasses() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM test_classes")
}
