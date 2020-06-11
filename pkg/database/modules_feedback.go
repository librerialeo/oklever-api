package database

import (
	"github.com/jackc/pgx"
)

// GetAllModulesFeedback queries for all modulesFeedback
func (db *Database) GetAllModulesFeedback() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM modules_feedback")
}
