package service

import (
	"github.com/jackc/pgx"
)

// GetAllSynchronousClassesResources return all synchronousClassesResources
func (s *Service) GetAllSynchronousClassesResources() (*pgx.Rows, error) {
	return s.db.GetAllSynchronousClassesResources()
}
