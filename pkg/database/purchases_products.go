package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllPurchasesProducts queries for all purchasesProducts
func (db *Database) GetAllPurchasesProducts() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM purchases_products")
}
