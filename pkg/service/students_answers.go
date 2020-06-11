package service

import (
	"github.com/jackc/pgx"
)

// GetAllStudentsAnswers return all studentsAnswers
func (s *Service) GetAllStudentsAnswers() (*pgx.Rows, error) {
	return s.db.GetAllStudentsAnswers()
}
