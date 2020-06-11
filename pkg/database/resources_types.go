package database

import (
	"github.com/jackc/pgx"
)

// GetAllResourcesTypes queries for all resourcesTypes
func (db *Database) GetAllResourcesTypes() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM resources_types")
}
