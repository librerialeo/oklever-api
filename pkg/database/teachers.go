package database

import (
	"github.com/jackc/pgx"
)

// GetAllTeachers queries for all teachers
func (db *Database) GetAllTeachers() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM teachers")
}

// GetTeacherByEmail get the user who's owns the pased email
func (db *Database) GetTeacherByEmail(email string) (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM users WHERE user_email = $1 AND rol_id = 2", email)
}

// GetTeacherByUserID get teacher information by user id
func (db *Database) GetTeacherByUserID(userID int32) (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM users WHERE user_id = $1", userID)
}

// UpdateTeacherInformation update teacher information
func (db *Database) UpdateTeacherInformation(userID int32, license string, rfc string) (*pgx.Rows, error) {
	return db.conn.Query("UPDATE users SET user_license=$1, user_rfc=$2 WHERE user_id=$3 RETURNING user_license, user_rfc", license, rfc, userID)
}
