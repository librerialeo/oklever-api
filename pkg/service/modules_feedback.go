package service

import (
	"github.com/jackc/pgx"
)

// GetAllModulesFeedback return all modulesFeedback
func (s *Service) GetAllModulesFeedback() (pgx.Rows, error) {
	return s.db.GetAllModulesFeedback()
}
