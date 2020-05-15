package database

import (
	"context"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx"
	"github.com/savsgio/atreugo"
)

// DBInvestment is user_research database table structure
type DBInvestment struct {
	ID        pgtype.Int4        `json:"id"`
	UserID    pgtype.Int4        `json:"user"`
	Reference pgtype.Text        `json:"reference"`
	Year      pgtype.Int4        `json:"year"`
	Type      pgtype.Varchar     `json:"type"`
	Added     pgtype.Timestamptz `json:"added"`
	Modified  pgtype.Timestamptz `json:"modified"`
	Deleted   pgtype.Timestamptz `json:"deleted"`
}

// ParseForUser as name says
func (investment *DBInvestment) ParseForUser() atreugo.JSON {
	return atreugo.JSON{
		"id":        investment.ID.Int,
		"reference": investment.Reference.String,
		"year":      investment.Year.Int,
		"type":      investment.Type.String,
	}
}

// GetUserInvestment get all userID  investments
func (db *Database) GetUserInvestment(investmentID int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users_research WHERE user_research_id=$1", investmentID)
}

// GetUserInvestments get all userID  investments
func (db *Database) GetUserInvestments(userID int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users_research WHERE user_id=$1", userID)
}

// AddUserInvestment get all userID  investments
func (db *Database) AddUserInvestment(userID int32, typeof string, reference string, year int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "INSERT INTO users_research (user_id, user_research_type, user_research_reference, user_research_year) values ($1, $2, $3, $4) RETURNING *", userID, typeof, reference, year)
}

// UpdateUserInvestment get all userID  investments
func (db *Database) UpdateUserInvestment(investmentID int32, reference string, year int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "UPDATE users_research SET user_research_reference=$1, user_research_year=$2 WHERE user_research_id=$3 RETURNING *", reference, year, investmentID)
}

// DeleteUserInvestment get all userID  investments
func (db *Database) DeleteUserInvestment(investmentID int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "DELETE FROM users_research WHERE user_research_id=$1", investmentID)
}
