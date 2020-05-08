package database

import (
	"context"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx"
)

// DBUsersAcademy struct of database users_languages
type DBUsersAcademy struct {
	ID                     pgtype.Int4    `json:"-"`
	UserID                 pgtype.Int4    `json:"-"`
	DegreeID               pgtype.Int4    `json:"degree"`
	UserAcademyName        pgtype.Varchar `json:"name"`
	UserAcademyInstitution pgtype.Varchar `json:"institution"`
	UserAcademyYear        pgtype.Int4    `json:"year"`
}

// GetAllUsersAcademy queries for all usersAcademy
func (db *Database) GetAllUsersAcademy(UserID int) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users_academy where user_id = $1", UserID)
}

// AddUsersAcademy add new user academy
func (db *Database) AddUsersAcademy(userID int32, degreeID int, academyName string, institution string, year int) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "", userID, degreeID, academyName, institution, year)
}
