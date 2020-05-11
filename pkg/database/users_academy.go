package database

import (
	"context"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx"
)

// DBUsersAcademy struct of database users_languages
type DBUsersAcademy struct {
	ID                     pgtype.Int4    `json:"id"`
	DegreeID               pgtype.Int4    `json:"degree"`
	UserAcademyName        pgtype.Varchar `json:"name"`
	UserAcademyInstitution pgtype.Varchar `json:"institution"`
	UserAcademyYear        pgtype.Int2    `json:"year"`
}

// GetAllUsersAcademy queries for all usersAcademy
func (db *Database) GetAllUsersAcademy(UserID int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT user_academy_id,degree_id,user_academy_name,user_academy_institution,user_academy_year FROM users_academy where user_id = $1", UserID)
}

// AddUsersAcademy add new user academy
func (db *Database) AddUsersAcademy(userID int32, degreeID int, academyName string, institution string, year int) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "INSERT INTO users_academy(user_id,degree_id,user_academy_name,user_academy_institution,user_academy_year) VALUES($1,$2,$3,$4,$5) RETURNING user_academy_id,degree_id,user_academy_name,user_academy_institution,user_academy_year", userID, degreeID, academyName, institution, year)
}
