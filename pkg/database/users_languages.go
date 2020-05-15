package database

import (
	"context"

	"github.com/jackc/pgtype"
	"github.com/savsgio/atreugo"

	"github.com/jackc/pgx"
)

// DBUsersLanguage struct of database users_languages
type DBUsersLanguage struct {
	ID         pgtype.Int4        `json:"-"`
	TeacherID  pgtype.Int4        `json:"_"`
	LanguageID pgtype.Int4        `json:"id"`
	CreatedAt  pgtype.Timestamptz `json:"-"`
	ModifiedAt pgtype.Timestamptz `json:"-"`
	DeleteAt   pgtype.Timestamptz `json:"-"`
}

// ParseForUser as name says
func (language *DBUsersLanguage) ParseForUser() atreugo.JSON {
	return atreugo.JSON{
		"id":       language.ID.Int,
		"language": language.LanguageID.Int,
	}
}

// GetAllUsersLanguages queries for all usersLanguages
func (db *Database) GetAllUsersLanguages(userID int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT language_id FROM users_languages where user_id = $1", userID)
}

// AddUsersLanguages add language users
func (db *Database) AddUsersLanguages(userID int32, languageID int) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "INSERT INTO users_languages(user_id,language_id) VALUES($1,$2)", userID, languageID)
}

//DeleteUsersLanguages delete language users
func (db *Database) DeleteUsersLanguages(userID int32, languageID int) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "DELETE FROM users_languages where user_id = $1 AND language_id = $2", userID, languageID)
}
