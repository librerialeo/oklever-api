package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllInsignias() (pgx.Rows, error) {
	return s.db.getAllInsignias()
}