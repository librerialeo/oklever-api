package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllStudentsQuizzes queries for all studentsQuizzes
func (db *Database) GetAllStudentsQuizzes() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM students_quizzes")
}
