package database

import (
	"github.com/jackc/pgx"
)

// GetAllClasses queries for all classes
func (db *Database) GetAllClasses() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM classes")
}
