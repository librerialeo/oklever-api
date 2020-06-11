package database

import (
	"github.com/jackc/pgx"
)

// GetAllUsersDegrees queries for all usersDegrees
func (db *Database) GetAllUsersDegrees() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM users_degrees")
}
