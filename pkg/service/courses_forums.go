package service

import (
	"github.com/jackc/pgx"
)

// GetAllCoursesForums return all coursesForums
func (s *Service) GetAllCoursesForums() (pgx.Rows, error) {
	return s.db.GetAllCoursesForums()
}
