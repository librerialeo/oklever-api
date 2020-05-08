package database

import (
	"context"

	"github.com/jackc/pgtype"

	"github.com/jackc/pgx"
)

// DBUsersLanguage struct of database teachers_languages
type DBUsersLanguage struct {
	ID         pgtype.Int4        `json:"-"`
	TeacherID  pgtype.Int4        `json:"teacher_id"`
	LanguageID pgtype.Int4        `json:"id"`
	CreatedAt  pgtype.Timestamptz `json:"-"`
	ModifiedAt pgtype.Timestamptz `json:"-"`
	DeleteAt   pgtype.Timestamptz `json:"-"`
}

// GetAllUsersLanguages queries for all teachersLanguages
func (db *Database) GetAllUsersLanguages() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users_languages")
}

// AddUsersLanguages add language teacher
func (db *Database) AddUsersLanguages(teacherID int, languageID int) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "INSERT INTO users_languages(user_id,language_id) VALUES($1,$2)", teacherID, languageID)
}

//DeleteUsersLanguages delete language user
func (db *Database) DeleteUsersLanguages(userID int, languageID int) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "DELETE FROM users_languages where user_id = $1 AND language_id = $2", userID, languageID)
}
