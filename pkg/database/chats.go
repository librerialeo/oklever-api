package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllChats queries for all chats
func (db *Database) GetAllChats() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM chats")
}
