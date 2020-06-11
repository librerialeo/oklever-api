package service

import (
	"github.com/jackc/pgx"
)

// GetAllStudentsSubscriptions return all studentsSubscriptions
func (s *Service) GetAllStudentsSubscriptions() (*pgx.Rows, error) {
	return s.db.GetAllStudentsSubscriptions()
}
