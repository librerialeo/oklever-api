package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllQuestionsTypes() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM questions_types")
}
