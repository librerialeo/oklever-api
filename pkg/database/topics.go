package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllTopics queries for all topics
func (db *Database) GetAllTopics() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM topics")
}
