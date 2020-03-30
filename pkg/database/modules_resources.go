package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllModulesResources() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM modules_resources")
}
