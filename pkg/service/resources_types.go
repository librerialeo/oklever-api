package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllResourcesTypes() (pgx.Rows, error) {
	return s.db.getAllResourcesTypes()
}