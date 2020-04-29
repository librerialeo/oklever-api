package service

import (
	"github.com/jackc/pgx"
)

// GetAllTeachers return all teachers
func (s *Service) GetAllTeachers() (pgx.Rows, error) {
	return s.db.GetAllTeachers()
}

// AddTeacher adds a new teacher to database
func (s *Service) AddTeacher(firstname string, lastname string, email string, password string) (pgx.Rows, error) {
	return s.db.AddUser(firstname, lastname, email, password, 2)
}
