package service

import (
	"github.com/jackc/pgx"
)

// GetAllStudentsQuizzes return all studentsQuizzes
func (s *Service) GetAllStudentsQuizzes() (*pgx.Rows, error) {
	return s.db.GetAllStudentsQuizzes()
}
