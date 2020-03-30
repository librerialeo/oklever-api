package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllPurchases queries for all purchases
func (db *Database) GetAllPurchases() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM purchases")
}
