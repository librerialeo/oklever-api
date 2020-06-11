package service

import (
	"github.com/jackc/pgx"
)

// GetAllUsersDegrees return all usersDegrees
func (s *Service) GetAllUsersDegrees() (*pgx.Rows, error) {
	return s.db.GetAllUsersDegrees()
}
