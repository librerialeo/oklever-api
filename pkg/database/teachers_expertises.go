package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllTeachersExpertises queries for all teachersExpertises
func (db *Database) GetAllTeachersExpertises() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM teachers_expertises")
}
