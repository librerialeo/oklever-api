package database

import (
	"github.com/jackc/pgx"
)

// GetAllModulesResources queries for all modulesResources
func (db *Database) GetAllModulesResources() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM modules_resources")
}
