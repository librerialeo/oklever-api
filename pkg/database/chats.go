package database

import (
	"github.com/jackc/pgx"
)

// GetAllChats queries for all chats
func (db *Database) GetAllChats() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM chats")
}
