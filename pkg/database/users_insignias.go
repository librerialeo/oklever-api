package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllUsersInsignias queries for all usersInsignias
func (db *Database) GetAllUsersInsignias() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users_insignias")
}
