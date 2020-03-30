package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllCountries() (pgx.Rows, error) {
	return s.db.getAllCountries()
}