package database

import (
	"github.com/jackc/pgx"
)

// GetAllQuestionsTypes queries for all questionsTypes
func (db *Database) GetAllQuestionsTypes() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM questions_types")
}
