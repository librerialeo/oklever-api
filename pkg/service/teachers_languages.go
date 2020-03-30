package service

import (
	"github.com/jackc/pgx"
)

// GetAllTeachersLanguages return all teachersLanguages
func (s *Service) GetAllTeachersLanguages() (pgx.Rows, error) {
	return s.db.GetAllTeachersLanguages()
}
