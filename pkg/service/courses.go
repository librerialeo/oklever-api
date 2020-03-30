package service

import (
	"github.com/jackc/pgx"
)

// GetAllCourses return all courses
func (s *Service) GetAllCourses() (pgx.Rows, error) {
	return s.db.GetAllCourses()
}
