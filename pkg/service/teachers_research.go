package service

import (
	"github.com/jackc/pgx"
)

// GetAllTeachersResearch return all teachersResearch
func (s *Service) GetAllTeachersResearch() (pgx.Rows, error) {
	return s.db.GetAllTeachersResearch()
}
