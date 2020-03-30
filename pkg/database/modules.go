package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllModules() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM modules")
}
