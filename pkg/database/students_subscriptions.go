package database

import (
	"github.com/jackc/pgx"
)

// GetAllStudentsSubscriptions queries for all studentsSubscriptions
func (db *Database) GetAllStudentsSubscriptions() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM students_subscriptions")
}
