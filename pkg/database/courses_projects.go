package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllCoursesProjects queries for all coursesProjects
func (db *Database) GetAllCoursesProjects() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM courses_projects")
}
