package service

import (
	"github.com/jackc/pgx"
)

// GetAllResourcesTypes return all resourcesTypes
func (s *Service) GetAllResourcesTypes() (*pgx.Rows, error) {
	return s.db.GetAllResourcesTypes()
}
