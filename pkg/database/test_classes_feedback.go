package database

import (
	"github.com/jackc/pgx"
)

// GetAllTestClassesFeedback queries for all testClassesFeedback
func (db *Database) GetAllTestClassesFeedback() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM test_classes_feedback")
}
