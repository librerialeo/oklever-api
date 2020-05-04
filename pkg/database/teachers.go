package database

import (
	"context"

	"github.com/jackc/pgx"
)

// GetAllTeachers queries for all teachers
func (db *Database) GetAllTeachers() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM teachers")
}

// GetTeacherByEmail get the user who's owns the pased email
func (db *Database) GetTeacherByEmail(email string) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users WHERE user_email = $1", email)
}
