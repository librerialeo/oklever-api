package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllModulesResources() (pgx.Rows, error) {
	return s.db.getAllModulesResources()
}