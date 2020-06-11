package service

import (
	"github.com/jackc/pgx"
)

// GetAllClassesResources return all ClassesResources
func (s *Service) GetAllClassesResources() (*pgx.Rows, error) {
	return s.db.GetAllClassesResources()
}
