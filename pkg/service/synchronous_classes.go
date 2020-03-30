package service

import (
	"github.com/jackc/pgx"
)

// GetAllSynchronousClasses return all synchronousClasses
func (s *Service) GetAllSynchronousClasses() (pgx.Rows, error) {
	return s.db.GetAllSynchronousClasses()
}
