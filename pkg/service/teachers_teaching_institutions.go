package service

import (
	"github.com/jackc/pgx"
)

// GetAllTeachersTeachingInstitutions return all teachersTeachingInstitutions
func (s *Service) GetAllTeachersTeachingInstitutions() (pgx.Rows, error) {
	return s.db.GetAllTeachersTeachingInstitutions()
}
