package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllSynchronousClassesResources() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM synchronous_classes_resources")
}
