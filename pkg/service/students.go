package service

import (
	"github.com/jackc/pgx"
)

// GetAllStudents return all students
func (s *Service) GetAllStudents() (*pgx.Rows, error) {
	return s.db.GetAllStudents()
}
