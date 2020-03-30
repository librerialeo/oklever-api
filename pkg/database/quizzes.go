package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllQuizzes queries for all quizzes
func (db *Database) GetAllQuizzes() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM quizzes")
}
