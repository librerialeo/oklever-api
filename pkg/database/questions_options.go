package database

import (
	"github.com/jackc/pgx"
)

// GetAllQuestionsOptions queries for all questionsOptions
func (db *Database) GetAllQuestionsOptions() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM questions_options")
}
