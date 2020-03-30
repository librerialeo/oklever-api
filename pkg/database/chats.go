package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllChats() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM chats")
}
