package database

import (
	"github.com/jackc/pgx"
)

// GetAllStudentsProjects queries for all studentsProjects
func (db *Database) GetAllStudentsProjects() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM students_projects")
}
