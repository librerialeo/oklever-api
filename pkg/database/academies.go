package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAcademyByEmail get the user who's owns the pased email
func (db *Database) GetAcademyByEmail(email string) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users WHERE user_email = $1 AND rol_id IN (3, 4)", email)
}
