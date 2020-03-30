package service

import (
	"github.com/jackc/pgx"
)

// GetAllTeachers return all teachers
func (s *Service) GetAllTeachers() (pgx.Rows, error) {
	return s.db.GetAllTeachers()
}
