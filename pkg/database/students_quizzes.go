package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllStudentsQuizzes() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM students_quizzes")
}
