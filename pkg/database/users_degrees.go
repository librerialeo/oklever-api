package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllUsersDegrees() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users_degrees")
}
