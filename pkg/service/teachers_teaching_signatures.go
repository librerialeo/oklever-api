package service

import (
	"github.com/jackc/pgx"
)

// GetAllTeachersTeachingSignatures return all teachersTeachingSignatures
func (s *Service) GetAllTeachersTeachingSignatures() (pgx.Rows, error) {
	return s.db.GetAllTeachersTeachingSignatures()
}
