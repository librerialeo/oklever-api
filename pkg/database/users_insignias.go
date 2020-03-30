package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllUsersInsignias() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users_insignias")
}
