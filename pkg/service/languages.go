package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllLanguages() (pgx.Rows, error) {
	return s.db.getAllLanguages()
}