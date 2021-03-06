package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllQuestionsOptions queries for all questionsOptions
func (db *Database) GetAllQuestionsOptions() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM questions_options")
}
