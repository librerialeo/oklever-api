package service

import (
	"github.com/jackc/pgx"
)

// GetAllClassesComments return all ClassesComments
func (s *Service) GetAllClassesComments() (*pgx.Rows, error) {
	return s.db.GetAllClassesComments()
}
