package service

import (
	"github.com/jackc/pgx"
)

// GetAllTopicsResources return all topicsResources
func (s *Service) GetAllTopicsResources() (pgx.Rows, error) {
	return s.db.GetAllTopicsResources()
}
