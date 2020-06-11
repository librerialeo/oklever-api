package database

import (
	"github.com/jackc/pgx"
)

// GetAllUsersChats queries for all usersChats
func (db *Database) GetAllUsersChats() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM users_chats")
}
