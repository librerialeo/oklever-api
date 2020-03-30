package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllTeachersTeachingInstitutions queries for all teachersTeachingInstitutions
func (db *Database) GetAllTeachersTeachingInstitutions() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM teachers_teaching_institutions")
}
