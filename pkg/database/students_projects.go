package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllStudentsProjects queries for all studentsProjects
func (db *Database) GetAllStudentsProjects() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM students_projects")
}
