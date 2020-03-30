package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllRoles() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM roles")
}
