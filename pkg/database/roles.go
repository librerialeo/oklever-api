package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllRoles queries for all roles
func (db *Database) GetAllRoles() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM roles")
}
