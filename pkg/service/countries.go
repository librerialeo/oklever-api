package service

import (
	"github.com/jackc/pgx"
)

// GetAllCountries return all countries
func (s *Service) GetAllCountries() (pgx.Rows, error) {
	return s.db.GetAllCountries()
}
