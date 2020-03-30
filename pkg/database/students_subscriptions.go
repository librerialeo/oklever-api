package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllStudentsSubscriptions() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM students_subscriptions")
}
