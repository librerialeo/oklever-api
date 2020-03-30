package service

import (
	"github.com/jackc/pgx"
)

// GetAllTeachersTeaching return all teachersTeaching
func (s *Service) GetAllTeachersTeaching() (pgx.Rows, error) {
	return s.db.GetAllTeachersTeaching()
}
