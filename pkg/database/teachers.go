package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllTeachers queries for all teachers
func (db *Database) GetAllTeachers() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM teachers")
}
