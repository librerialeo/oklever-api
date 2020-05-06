package database

import (
	"github.com/jackc/pgx"
)

// Database is a wrapper for database actions
type Database struct {
	conn *pgx.Conn
}

// InitDatabase initialize database wrapper
func InitDatabase(conn *pgx.Conn) *Database {
	db := &Database{conn}
	return db
}
