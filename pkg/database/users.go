package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllUsers queries for all users
func (db *Database) GetAllUsers() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users")
}
