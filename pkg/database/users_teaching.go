package database

import (
	"context"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx"
)

// DBSignature is user_teaching_signatures database table structure
type DBSignature struct {
	ID       pgtype.Int4        `json:"id"`
	UserID   pgtype.Int4        `json:"user"`
	DegreeID pgtype.Int4        `json:"degree"`
	Name     pgtype.Varchar     `json:"name"`
	Added    pgtype.Timestamptz `json:"added"`
	Modified pgtype.Timestamptz `json:"modified"`
	Deleted  pgtype.Timestamptz `json:"deleted"`
}

// GetUserTeachingSignature get all userID teaching signatures
func (db *Database) GetUserTeachingSignature(signatureID int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users_teaching_signatures WHERE user_teaching_signature_id=$1", signatureID)
}

// GetUserTeachingSignatures get all userID teaching signatures
func (db *Database) GetUserTeachingSignatures(userID int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users_teaching WHERE user_id=$1", userID)
}

// AddUserTeachingSignature get all userID teaching signatures
func (db *Database) AddUserTeachingSignature(userID int32, degreeID int32, name string) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "INSERT INTO users_teaching_signatures (user_id, degree_id, user_teaching_signature_name) values ($1, $2, $3) RETURNING *", userID, degreeID, name)
}

// UpdateUserTeachingSignature get all userID teaching signatures
func (db *Database) UpdateUserTeachingSignature(signatureID int32, name string) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "UPDATE users_teaching_signatures SET user_teaching_signature_name=$1 WHERE user_teaching_signature_id=$2 RETURNING *", name, signatureID)
}

// DeleteUserTeachingSignature get all userID teaching signatures
func (db *Database) DeleteUserTeachingSignature(signatureID int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "DELETE FROM users_teaching_signatures WHERE user_teaching_signature_id=$1", signatureID)
}

// DBInstitution is user_teaching_institutions database table structure
type DBInstitution struct {
	ID       pgtype.Int4        `json:"id"`
	UserID   pgtype.Int4        `json:"user"`
	Name     pgtype.Varchar     `json:"name"`
	Added    pgtype.Timestamptz `json:"added"`
	Modified pgtype.Timestamptz `json:"modified"`
	Deleted  pgtype.Timestamptz `json:"deleted"`
}

// GetUserTeachingInstitution get all userID teaching institutions
func (db *Database) GetUserTeachingInstitution(institutionID int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users_teaching_institutions WHERE user_teaching_institution_id=$1", institutionID)
}

// GetUserTeachingInstitutions get all userID teaching institutions
func (db *Database) GetUserTeachingInstitutions(userID int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users_teaching_institutions WHERE user_id=$1", userID)
}

// AddUserTeachingInstitution get all userID teaching institutions
func (db *Database) AddUserTeachingInstitution(userID int32, name string) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "INSERT INTO users_teaching_institutions (user_id, user_teaching_institution_name) values ($1, $2) RETURNING *", userID, name)
}

// UpdateUserTeachingInstitution get all userID teaching institutions
func (db *Database) UpdateUserTeachingInstitution(institutionID int32, name string) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "UPDATE users_teaching_institutions SET user_teaching_institution_name=$1 WHERE user_teaching_institution_id=$2 RETURNING *", name, institutionID)
}

// DeleteUserTeachingInstitution get all userID teaching institutions
func (db *Database) DeleteUserTeachingInstitution(institutionID int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "DELETE FROM users_teaching_institutions WHERE user_teaching_institution_id=$1", institutionID)
}

// GetUserExperience get user teaching experience months
func (db *Database) GetUserExperience(userID int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT user_teaching_months FROM users WHERE user_id = $1", userID)
}

// SetUserExperience get user teaching experience months
func (db *Database) SetUserExperience(userID int32, months int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "UPDATE users SET user_teaching_months = $1 WHERE user_id = $2 RETURNING user_teaching_months", months, userID)
}
