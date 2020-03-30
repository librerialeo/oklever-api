package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllInsignias() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM insignias")
}
