package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllTeachersTeachingSignatures queries for all teachersTeachingSignatures
func (db *Database) GetAllTeachersTeachingSignatures() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM teachers_teaching_signatures")
}
