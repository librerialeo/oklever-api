package service

import (
	"github.com/jackc/pgx"
)

// GetAllTeachersExpertises return all teachersExpertises
func (s *Service) GetAllTeachersExpertises() (pgx.Rows, error) {
	return s.db.GetAllTeachersExpertises()
}
