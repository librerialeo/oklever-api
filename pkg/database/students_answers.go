package database

import (
	"github.com/jackc/pgx"
)

// GetAllStudentsAnswers queries for all studentsAnswers
func (db *Database) GetAllStudentsAnswers() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM students_answers")
}
