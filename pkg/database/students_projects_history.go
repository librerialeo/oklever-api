package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllStudentsProjectsHistory queries for all studentsProjectsHistory
func (db *Database) GetAllStudentsProjectsHistory() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM students_projects_history")
}
