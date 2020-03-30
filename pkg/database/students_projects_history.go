package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllStudentsProjectsHistory() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM students_projects_history")
}
