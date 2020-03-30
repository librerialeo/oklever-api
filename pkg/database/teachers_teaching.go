package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllTeachersTeaching queries for all teachersTeaching
func (db *Database) GetAllTeachersTeaching() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM teachers_teaching")
}
