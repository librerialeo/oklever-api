package database

import (
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx"
	"github.com/savsgio/atreugo/v11"
)

// DBUsersAcademy struct of database users_languages
type DBUsersAcademy struct {
	ID                     pgtype.Int4    `json:"id"`
	DegreeID               pgtype.Int4    `json:"degree"`
	UserAcademyName        pgtype.Varchar `json:"name"`
	UserAcademyInstitution pgtype.Varchar `json:"institution"`
	UserAcademyYear        pgtype.Int2    `json:"year"`
}

// ParseForUser as name says
func (academy *DBUsersAcademy) ParseForUser() atreugo.JSON {
	return atreugo.JSON{
		"id":          academy.ID.Int,
		"degree":      academy.DegreeID.Int,
		"name":        academy.UserAcademyName.String,
		"institution": academy.UserAcademyInstitution.String,
		"year":        academy.UserAcademyYear.Int,
	}
}

// GetUserAcademiesByUserID queries for all usersAcademy
func (db *Database) GetUserAcademiesByUserID(UserID int32) (*pgx.Rows, error) {
	return db.conn.Query("SELECT user_academy_id,degree_id,user_academy_name,user_academy_institution,user_academy_year FROM users_academy where user_id = $1", UserID)
}

// AddUsersAcademy add new user academy
func (db *Database) AddUsersAcademy(userID int32, degreeID int, academyName string, institution string, year int) (*pgx.Rows, error) {
	return db.conn.Query("INSERT INTO users_academy(user_id,degree_id,user_academy_name,user_academy_institution,user_academy_year) VALUES($1,$2,$3,$4,$5) RETURNING user_academy_id,degree_id,user_academy_name,user_academy_institution,user_academy_year", userID, degreeID, academyName, institution, year)
}

// UpdateUsersAcademy add new user academy
func (db *Database) UpdateUsersAcademy(ID int, userID int32, degreeID int, academyName string, institution string, year int) (*pgx.Rows, error) {
	return db.conn.Query("UPDATE users_academy SET degree_id = $1, user_academy_name = $2, user_academy_institution = $3, user_academy_year = $4 WHERE user_id = $6 AND user_academy_id = $5 RETURNING user_academy_id,degree_id,user_academy_name,user_academy_institution,user_academy_year", degreeID, academyName, institution, year, ID, userID)
}

// DeleteUsersAcademy add new user academy
func (db *Database) DeleteUsersAcademy(ID int, userID int32) (*pgx.Rows, error) {
	return db.conn.Query("DELETE FROM users_academy WHERE user_academy_id = $1 AND user_id = $2", ID, userID)
}
