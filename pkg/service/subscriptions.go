package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllSubscriptions() (pgx.Rows, error) {
	return s.db.getAllSubscriptions()
}