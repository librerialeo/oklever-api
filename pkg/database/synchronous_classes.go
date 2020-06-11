package database

import (
	"github.com/jackc/pgx"
)

// GetAllSynchronousClasses queries for all synchronousClasses
func (db *Database) GetAllSynchronousClasses() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM synchronous_classes")
}
