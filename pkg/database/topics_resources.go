package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllTopicsResources() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM topics_resources")
}
