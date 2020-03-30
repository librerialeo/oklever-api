package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllQuestionsTypes queries for all questionsTypes
func (db *Database) GetAllQuestionsTypes() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM questions_types")
}
