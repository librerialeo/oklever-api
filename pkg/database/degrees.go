package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllDegrees() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM degrees")
}
