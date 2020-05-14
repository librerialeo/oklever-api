package database

import (
	"context"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx"
)

// DBManagement is user_management database table structure
type DBManagement struct {
	ID          pgtype.Int4        `json:"id"`
	UserID      pgtype.Int4        `json:"user"`
	Job         pgtype.Varchar     `json:"job"`
	Institution pgtype.Varchar     `json:"institution"`
	Months      pgtype.Int2        `json:"months"`
	Added       pgtype.Timestamptz `json:"added"`
	Modified    pgtype.Timestamptz `json:"modified"`
	Deleted     pgtype.Timestamptz `json:"deleted"`
}

// GetUserManagement get all userID  managements
func (db *Database) GetUserManagement(managementID int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users_managements WHERE user_management_id=$1", managementID)
}

// GetUserManagements get all userID  managements
func (db *Database) GetUserManagements(userID int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users_managements WHERE user_id=$1", userID)
}

// AddUserManagement get all userID  managements
func (db *Database) AddUserManagement(userID int32, job string, institution string, months int16) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "INSERT INTO users_managements (user_id, user_management_job, user_management_institution, user_management_months) values ($1, $2, $3, $4) RETURNING *", userID, job, institution, months)
}

// UpdateUserManagement get all userID  managements
func (db *Database) UpdateUserManagement(managementID int32, job string, institution string, months int16) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "UPDATE users_managements SET user_management_job=$1, user_management_institution=$2, user_management_months=$3 WHERE user_management_id=$4 RETURNING *", job, institution, months, managementID)
}

// DeleteUserManagement get all userID  managements
func (db *Database) DeleteUserManagement(managementID int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "DELETE FROM users_managements WHERE user_management_id=$1", managementID)
}
