package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllResourcesTypes() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM resources_types")
}
