package database

import (
	"context"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx"
)

// DBExpertise is user_expertise database table structure
type DBExpertise struct {
	ID       pgtype.Int4        `json:"id"`
	UserID   pgtype.Int4        `json:"user"`
	Name     pgtype.Varchar     `json:"name"`
	Months   pgtype.Int2        `json:"months"`
	Added    pgtype.Timestamptz `json:"added"`
	Modified pgtype.Timestamptz `json:"modified"`
	Deleted  pgtype.Timestamptz `json:"deleted"`
}

// GetUserExpertise get all userID  expertises
func (db *Database) GetUserExpertise(expertiseID int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users_expertises WHERE user_expertise_id=$1", expertiseID)
}

// GetUserExpertises get all userID  expertises
func (db *Database) GetUserExpertises(userID int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users_expertises WHERE user_id=$1", userID)
}

// AddUserExpertise get all userID  expertises
func (db *Database) AddUserExpertise(userID int32, name string, months int16) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "INSERT INTO users_expertises (user_id, user_expertise_name, user_expertise_months) values ($1, $2, $3) RETURNING *", userID, name, months)
}

// UpdateUserExpertise get all userID  expertises
func (db *Database) UpdateUserExpertise(expertiseID int32, name string, months int16) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "UPDATE users_expertises SET user_expertise_name=$1, user_expertise_months=$2 WHERE user_expertise_id=$3 RETURNING *", name, months, expertiseID)
}

// DeleteUserExpertise get all userID  expertises
func (db *Database) DeleteUserExpertise(expertiseID int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "DELETE FROM users_expertises WHERE user_expertise_id=$1", expertiseID)
}
