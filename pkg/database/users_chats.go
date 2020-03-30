package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllUsersChats queries for all usersChats
func (db *Database) GetAllUsersChats() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users_chats")
}
