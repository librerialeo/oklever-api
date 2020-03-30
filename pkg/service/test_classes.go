package service

import (
	"github.com/jackc/pgx"
)

// GetAllTestClasses return all testClasses
func (s *Service) GetAllTestClasses() (pgx.Rows, error) {
	return s.db.GetAllTestClasses()
}
