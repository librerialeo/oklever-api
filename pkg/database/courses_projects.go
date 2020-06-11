package database

import (
	"github.com/jackc/pgx"
)

// GetAllCoursesProjects queries for all coursesProjects
func (db *Database) GetAllCoursesProjects() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM courses_projects")
}
