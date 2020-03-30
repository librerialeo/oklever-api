package service

import (
	"github.com/jackc/pgx"
)

// GetAllStudentsCourses return all studentsCourses
func (s *Service) GetAllStudentsCourses() (pgx.Rows, error) {
	return s.db.GetAllStudentsCourses()
}
