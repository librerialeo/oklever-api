package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllTeachersLanguages queries for all teachersLanguages
func (db *Database) GetAllTeachersLanguages() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM teachers_languages")
}
