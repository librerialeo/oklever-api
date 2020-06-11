package database

import (
	"github.com/jackc/pgx"
)

// GetAllTopics queries for all topics
func (db *Database) GetAllTopics() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM topics")
}
