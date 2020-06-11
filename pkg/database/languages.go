package database

import (
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx"
)

// DBLanguage struct of database
type DBLanguage struct {
	ID       pgtype.Int4      `json:"id"`
	Name     pgtype.Varchar   `json:"name"`
	ISO      pgtype.Varchar   `json:"iso"`
	Created  pgtype.Timestamp `json:"created"`
	Modified pgtype.Timestamp `json:"modifed"`
	Deleted  pgtype.Timestamp `json:"deleted"`
}

// GetAllLanguages queries for all languages
func (db *Database) GetAllLanguages() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM languages")
}
