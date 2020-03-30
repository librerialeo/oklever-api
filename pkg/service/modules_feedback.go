package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllModulesFeedback() (pgx.Rows, error) {
	return s.db.getAllModulesFeedback()
}