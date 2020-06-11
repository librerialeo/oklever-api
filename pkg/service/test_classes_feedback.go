package service

import (
	"github.com/jackc/pgx"
)

// GetAllTestClassesFeedback return all testClassesFeedback
func (s *Service) GetAllTestClassesFeedback() (*pgx.Rows, error) {
	return s.db.GetAllTestClassesFeedback()
}
