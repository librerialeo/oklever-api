package service

import (
	"github.com/jackc/pgx"
)

// GetAllModulesResources return all modulesResources
func (s *Service) GetAllModulesResources() (pgx.Rows, error) {
	return s.db.GetAllModulesResources()
}
