package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllQuizzes() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM quizzes")
}
