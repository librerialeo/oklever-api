package service

import (
	"github.com/jackc/pgx"
)

// GetAllQuizzes return all quizzes
func (s *Service) GetAllQuizzes() (*pgx.Rows, error) {
	return s.db.GetAllQuizzes()
}
