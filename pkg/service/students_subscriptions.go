package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllStudentsSubscriptions() (pgx.Rows, error) {
	return s.db.getAllStudentsSubscriptions()
}