package database

import (
	"github.com/jackc/pgx"
)

// GetAllStudentsQuizzes queries for all studentsQuizzes
func (db *Database) GetAllStudentsQuizzes() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM students_quizzes")
}
