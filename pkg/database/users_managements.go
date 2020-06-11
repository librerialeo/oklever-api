package database

import (
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx"
	"github.com/savsgio/atreugo/v11"
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

// ParseForUser as name says
func (management *DBManagement) ParseForUser() atreugo.JSON {
	return atreugo.JSON{
		"id":          management.ID.Int,
		"job":         management.Job.String,
		"institution": management.Institution.String,
		"months":      management.Months.Int,
	}
}

// GetUserManagement get all userID  managements
func (db *Database) GetUserManagement(managementID int32) (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM users_managements WHERE user_management_id=$1", managementID)
}

// GetUserManagements get all userID  managements
func (db *Database) GetUserManagements(userID int32) (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM users_managements WHERE user_id=$1", userID)
}

// AddUserManagement get all userID  managements
func (db *Database) AddUserManagement(userID int32, job string, institution string, months int16) (*pgx.Rows, error) {
	return db.conn.Query("INSERT INTO users_managements (user_id, user_management_job, user_management_institution, user_management_months) values ($1, $2, $3, $4) RETURNING *", userID, job, institution, months)
}

// UpdateUserManagement get all userID  managements
func (db *Database) UpdateUserManagement(managementID int32, job string, institution string, months int16) (*pgx.Rows, error) {
	return db.conn.Query("UPDATE users_managements SET user_management_job=$1, user_management_institution=$2, user_management_months=$3 WHERE user_management_id=$4 RETURNING *", job, institution, months, managementID)
}

// DeleteUserManagement get all userID  managements
func (db *Database) DeleteUserManagement(managementID int32) (*pgx.Rows, error) {
	return db.conn.Query("DELETE FROM users_managements WHERE user_management_id=$1", managementID)
}
