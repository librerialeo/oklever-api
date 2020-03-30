package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllTeachersManagements queries for all teachersManagements
func (db *Database) GetAllTeachersManagements() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM teachers_managements")
}
