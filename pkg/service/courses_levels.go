package service

import (
	"github.com/jackc/pgx"
)

// GetAllCoursesLevels return all coursesLevels
func (s *Service) GetAllCoursesLevels() (*pgx.Rows, error) {
	return s.db.GetAllCoursesLevels()
}
