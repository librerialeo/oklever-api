package service

import (
	"github.com/jackc/pgx"
)

// GetAllDegrees return all degrees
func (s *Service) GetAllDegrees() (pgx.Rows, error) {
	return s.db.GetAllDegrees()
}
