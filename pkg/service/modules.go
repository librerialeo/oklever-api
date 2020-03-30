package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllModules() (pgx.Rows, error) {
	return s.db.getAllModules()
}