package service

import (
	"github.com/jackc/pgx"
)

// GetAllQuestionsOptions return all questionsOptions
func (s *Service) GetAllQuestionsOptions() (*pgx.Rows, error) {
	return s.db.GetAllQuestionsOptions()
}
