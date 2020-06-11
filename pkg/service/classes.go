package service

import (
	"github.com/jackc/pgx"
)

// GetAllClasses return all classes
func (s *Service) GetAllClasses() (*pgx.Rows, error) {
	return s.db.GetAllClasses()
}
