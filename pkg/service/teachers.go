package service

import (
	"errors"

	"github.com/jackc/pgx"
	"github.com/librerialeo/oklever-api/pkg/database"
)

// GetAllTeachers return all teachers
func (s *Service) GetAllTeachers() (pgx.Rows, error) {
	return s.db.GetAllTeachers()
}

// AddTeacher adds a new teacher to database
func (s *Service) AddTeacher(firstname string, lastname string, email string, password string) (*database.DBUser, error) {
	rows, err := s.db.AddUser(firstname, lastname, email, password, 2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var u database.DBUser
	if rows.Next() {
		err = rows.Scan(&u.ID,
			&u.Email,
			&u.Password,
			&u.Firstname,
			&u.Lastname,
			&u.Gender,
			&u.Image,
			&u.Birthdate,
			&u.Phone,
			&u.Country,
			&u.Rol,
			&u.LastAction,
			&u.Created,
			&u.Modified,
			&u.Deleted)
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// GetTeacherUserByEmail gets the teacher that owns email
func (s *Service) GetTeacherUserByEmail(email string) (*database.DBUser, error) {
	rows, err := s.db.GetTeacherByEmail(email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var u database.DBUser
	if rows.CommandTag().RowsAffected() > 1 {
		return nil, errors.New("dupplicated email")
	} else if rows.CommandTag().RowsAffected() > 1 {
		return nil, errors.New("user not found")
	}
	if rows.Next() {
		err = rows.Scan(&u.ID,
			&u.Email,
			&u.Password,
			&u.Firstname,
			&u.Lastname,
			&u.Gender,
			&u.Image,
			&u.Birthdate,
			&u.Phone,
			&u.Country,
			&u.Rol,
			&u.LastAction,
			&u.Created,
			&u.Modified,
			&u.Deleted)
	}
	return &u, err
}
