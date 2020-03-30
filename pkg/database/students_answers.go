package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllStudentsAnswers() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM students_answers")
}
