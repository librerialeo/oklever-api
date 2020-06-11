package database

import (
	"github.com/jackc/pgx"
)

// GetAllPurchasesProducts queries for all purchasesProducts
func (db *Database) GetAllPurchasesProducts() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM purchases_products")
}
