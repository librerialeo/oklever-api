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
	return db.conn.Query(context.Background(), "SELECT * FROM users WHERE user_email = $1 AND rol_id = 2", email)
}

// GetTeacherByUserID get teacher information by user id
func (db *Database) GetTeacherByUserID(userID int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM teachers WHERE user_id = $1")
}

// UpdateTeacherInformation update teacher information
func (db *Database) UpdateTeacherInformation(userID int32, license string, rfc string) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "UPDATE users SET user_license=$1, teacher_rfc=$2 WHERE user_id=$3", license, rfc, userID)
}
