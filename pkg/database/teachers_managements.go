package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllTeachersManagements() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM teachers_managements")
}
