package service

import (
	"github.com/jackc/pgx"
)

// GetAllQuestions return all questions
func (s *Service) GetAllQuestions() (*pgx.Rows, error) {
	return s.db.GetAllQuestions()
}
