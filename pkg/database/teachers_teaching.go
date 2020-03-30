package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllTeachersTeaching() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM teachers_teaching")
}
