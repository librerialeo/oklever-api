package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllModulesFeedback queries for all modulesFeedback
func (db *Database) GetAllModulesFeedback() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM modules_feedback")
}
