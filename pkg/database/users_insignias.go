package database

import (
	"github.com/jackc/pgx"
)

// GetAllUsersInsignias queries for all usersInsignias
func (db *Database) GetAllUsersInsignias() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM users_insignias")
}
