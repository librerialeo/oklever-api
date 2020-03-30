package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllTeachersResearch queries for all teachersResearch
func (db *Database) GetAllTeachersResearch() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM teachers_research")
}
