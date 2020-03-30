package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllModulesFeedback() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM modules_feedback")
}
