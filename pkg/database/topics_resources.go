package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllTopicsResources queries for all topicsResources
func (db *Database) GetAllTopicsResources() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM topics_resources")
}
