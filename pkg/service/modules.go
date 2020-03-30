package service

import (
	"github.com/jackc/pgx"
)

// GetAllModules return all modules
func (s *Service) GetAllModules() (pgx.Rows, error) {
	return s.db.GetAllModules()
}
