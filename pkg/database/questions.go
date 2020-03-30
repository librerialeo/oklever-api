package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllQuestions queries for all questions
func (db *Database) GetAllQuestions() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM questions")
}
