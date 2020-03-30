package database

import (
	"context"

	"github.com/jackc/pgx"
)

func (db *Database) getAllPurchasesProducts() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM purchases_products")
}
