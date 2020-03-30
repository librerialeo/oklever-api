package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllUsersDegrees queries for all usersDegrees
func (db *Database) GetAllUsersDegrees() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users_degrees")
}
