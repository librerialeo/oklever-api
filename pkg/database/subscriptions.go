package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllSubscriptions queries for all subscriptions
func (db *Database) GetAllSubscriptions() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM subscriptions")
}
