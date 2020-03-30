package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllUsersChats() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users_chats")
}
