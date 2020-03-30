package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllStudentsSubscriptions queries for all studentsSubscriptions
func (db *Database) GetAllStudentsSubscriptions() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM students_subscriptions")
}
