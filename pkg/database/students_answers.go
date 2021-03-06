package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllStudentsAnswers queries for all studentsAnswers
func (db *Database) GetAllStudentsAnswers() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM students_answers")
}
