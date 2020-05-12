package database

import (
	"context"

	"github.com/jackc/pgtype"

	"github.com/jackc/pgx"
)

//DBDegrees struct of database degrees
type DBDegrees struct {
	ID          pgtype.Int4    `json:"id"`
	Name        pgtype.Varchar `json:"name"`
	DegreeClass pgtype.Varchar `json:"class"`
}

// GetAllDegrees queries for all degrees
func (db *Database) GetAllDegrees() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT degree_id, degree_name, degree_class FROM degrees")
}
