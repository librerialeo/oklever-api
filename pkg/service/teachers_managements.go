package service

import (
	"github.com/jackc/pgx"
)

// GetAllTeachersManagements return all teachersManagements
func (s *Service) GetAllTeachersManagements() (pgx.Rows, error) {
	return s.db.GetAllTeachersManagements()
}
