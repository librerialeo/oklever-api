package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllResourcesTypes queries for all resourcesTypes
func (db *Database) GetAllResourcesTypes() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM resources_types")
}
