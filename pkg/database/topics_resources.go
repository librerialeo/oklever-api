package database

import (
	"github.com/jackc/pgx"
)

// GetAllTopicsResources queries for all topicsResources
func (db *Database) GetAllTopicsResources() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM topics_resources")
}
