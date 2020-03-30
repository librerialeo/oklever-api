package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllModulesResources queries for all modulesResources
func (db *Database) GetAllModulesResources() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM modules_resources")
}
