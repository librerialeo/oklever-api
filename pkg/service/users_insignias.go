package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllUsersInsignias() (pgx.Rows, error) {
	return s.db.getAllUsersInsignias()
}