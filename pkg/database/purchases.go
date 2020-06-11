package database

import (
	"github.com/jackc/pgx"
)

// GetAllPurchases queries for all purchases
func (db *Database) GetAllPurchases() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM purchases")
}
