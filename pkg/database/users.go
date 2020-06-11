package database

import (
	"time"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx"
	"github.com/savsgio/atreugo/v11"
)

// DBUser struct of database user
type DBUser struct {
	ID             pgtype.Int4        `json:"id"`
	Email          pgtype.Varchar     `json:"email"`
	Password       pgtype.Varchar     `json:"-"`
	Firstname      pgtype.Varchar     `json:"firstname"`
	Lastname       pgtype.Varchar     `json:"lastname"`
	Gender         pgtype.Varchar     `json:"gender"`
	Image          pgtype.Varchar     `json:"image"`
	Birthdate      pgtype.Date        `json:"birthdate"`
	Phone          pgtype.Varchar     `json:"phone"`
	License        pgtype.Varchar     `json:"license"`
	RFC            pgtype.Varchar     `json:"rfc"`
	Biography      pgtype.Text        `json:"biography"`
	TeachingMonths pgtype.Int2        `json:"teaching_months"`
	Status         pgtype.Varchar     `json:"status"`
	Country        pgtype.Int4        `json:"country"`
	Rol            pgtype.Int4        `json:"rol"`
	LastAction     pgtype.Timestamptz `json:"last_action"`
	Created        pgtype.Timestamptz `json:"created"`
	Modified       pgtype.Timestamptz `json:"modified"`
	Deleted        pgtype.Timestamptz `json:"deleted"`
}

// ParseForUser parse DBuser for users
func (user *DBUser) ParseForUser() atreugo.JSON {
	return atreugo.JSON{
		"gender":          user.Gender.String,
		"image":           user.Image.String,
		"birthdate":       user.Birthdate.Time,
		"phone":           user.Phone.String,
		"license":         user.License.String,
		"rfc":             user.RFC.String,
		"biography":       user.Biography.String,
		"teaching_months": user.TeachingMonths.Int,
		"status":          user.Status.String,
		"country":         user.Country.Int,
	}
}

// GetAllUsers queries for all users
func (db *Database) GetAllUsers() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM users")
}

// AddUser adds new user
func (db *Database) AddUser(firstname string, lastname string, email string, password string, rol int32) (*pgx.Rows, error) {
	return db.conn.Query("INSERT INTO users(user_firstname,user_lastname,user_email,user_password,rol_id) VALUES($1,$2,$3,$4,$5) RETURNING user_id, user_firstname, user_lastname, user_email, rol_id, user_status", firstname, lastname, email, password, rol)
}

// GetUserLastAction get the user last action from database
func (db *Database) GetUserLastAction(userID int32) (*pgx.Rows, error) {
	return db.conn.Query("SELECT user_lastaction FROM users WHERE user_id = $1", userID)
}

// UpdateUserLastAction sets the user las action in database
func (db *Database) UpdateUserLastAction(userID int32, lastaction time.Time) (*pgx.Rows, error) {
	return db.conn.Query("UPDATE users SET user_lastaction = $1 WHERE user_id = $2 RETURNING user_lastaction", lastaction, userID)
}

// UpdateUserInformation update user information
func (db *Database) UpdateUserInformation(userID int32, first string, last string, email string, gender string, phone string) (*pgx.Rows, error) {
	return db.conn.Query("UPDATE users SET user_firstname=$1, user_lastname=$2, user_email=$3, user_gender=$4, user_phone=$5 WHERE user_id=$6", first, last, email, gender, phone, userID)
}

// GetUserBiography get user biography
func (db *Database) GetUserBiography(userID int32) (*pgx.Rows, error) {
	return db.conn.Query("SELECT user_biography FROM users WHERE user_id = $1 and deleted_at IS NULL", userID)
}

// SetUserBiography get user biography
func (db *Database) SetUserBiography(userID int32, biography string) (*pgx.Rows, error) {
	return db.conn.Query("UPDATE users SET user_biography=$1 WHERE user_id = $2 RETURNING user_biography", biography, userID)
}

// SetUserImage get user biography
func (db *Database) SetUserImage(userID int32, image string) (*pgx.Rows, error) {
	return db.conn.Query("UPDATE users SET user_image=$1 WHERE user_id = $2 RETURNING user_image", image, userID)
}

// SetUserStatus set user status
func (db *Database) SetUserStatus(userID int32, status string) (*pgx.Rows, error) {
	return db.conn.Query("UPDATE users SET user_status = $1 WHERE user_id = $2 RETURNING user_image", status, userID)
}

// GetUserByID gets user with passed id from database if exits
func (db *Database) GetUserByID(userID int32) (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM users WHERE user_id = $1", userID)
}

// GetUserCredentialsByID gets user with passed id from database if exits
func (db *Database) GetUserCredentialsByID(userID int32) (*pgx.Rows, error) {
	return db.conn.Query("SELECT user_id, rol_id, user_status FROM users WHERE user_id = $1", userID)
}
