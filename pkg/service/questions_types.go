package service

import (
	"github.com/jackc/pgx"
)

// GetAllQuestionsTypes return all questionsTypes
func (s *Service) GetAllQuestionsTypes() (*pgx.Rows, error) {
	return s.db.GetAllQuestionsTypes()
}
