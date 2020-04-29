package database

import (
	"context"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx"
)

// DBUser struct of database user
type DBUser struct {
	ID        pgtype.XID
	Email     pgtype.Varchar
	Password  pgtype.Varchar
	Firstname pgtype.Varchar
	Lastname  pgtype.Varchar
	Gender    pgtype.Varchar
	Image     pgtype.Varchar
	Birthdate pgtype.Date
	Phone     pgtype.Varchar
	Country   pgtype.Int4
	Rol       pgtype.Int4
	Created   pgtype.Timestamptz
	Modified  pgtype.Timestamptz
	Deleted   pgtype.Timestamptz
}

// GetAllUsers queries for all users
func (db *Database) GetAllUsers() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM users")
}

// AddUser adds new user
func (db *Database) AddUser(firstname string, lastname string, email string, password string, rol int) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "addUser", firstname, lastname, email, password, rol)
	// return db.conn.Query(context.Background(),
	// 	"INSERT INTO users(user_firstname,user_lastname,user_email,user_password,rol_id) VALUES($1,$2,$3,$4,$5)",
	// 	firstname, lastname, email, password, rol)
}

// InitUserQueries initialize all user quieries
func (db *Database) InitUserQueries() {
	db.conn.Prepare(context.Background(), "addUser", "INSERT INTO users(user_firstname,user_lastname,user_email,user_password,rol_id) VALUES($1,$2,$3,$4,$5) RETURNING user_id, user_email, user_firstname, user_lastname, rol_id, created_at")
}
