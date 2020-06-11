package database

import (
	"github.com/jackc/pgx"
)

// GetAllStudentsProjectsHistory queries for all studentsProjectsHistory
func (db *Database) GetAllStudentsProjectsHistory() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM students_projects_history")
}
