package database

import (
	"context"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx"
)

// DBUser struct of database user
type DBUser struct {
	ID        pgtype.Int4        `json:"-"`
	Email     pgtype.Varchar     `json:"email"`
	Password  pgtype.Varchar     `json:"-"`
	Firstname pgtype.Varchar     `json:"firstname"`
	Lastname  pgtype.Varchar     `json:"lastname"`
	Gender    pgtype.Varchar     `json:"gender"`
	Image     pgtype.Varchar     `json:"image"`
	Birthdate pgtype.Date        `json:"birthdate"`
	Phone     pgtype.Varchar     `json:"phone"`
	Country   pgtype.Int4        `json:"country"`
	Rol       pgtype.Int4        `json:"-"`
	Created   pgtype.Timestamptz `json:"-"`
	Modified  pgtype.Timestamptz `json:"-"`
	Deleted   pgtype.Timestamptz `json:"-"`
}

// GetAllUsers queries for all users
func (db *Database) GetAllUsers() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users")
}

// AddUser adds new user
func (db *Database) AddUser(firstname string, lastname string, email string, password string, rol int) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "addUser", firstname, lastname, email, password, rol)
}

// InitUserQueries initialize all user quieries
func (db *Database) InitUserQueries() {
	db.conn.Prepare(context.Background(), "addUser", "INSERT INTO users(user_firstname,user_lastname,user_email,user_password,rol_id) VALUES($1,$2,$3,$4,$5) RETURNING *")
}
