package service

import (
	"github.com/jackc/pgx"
)

// GetAllCoursesProjects return all coursesProjects
func (s *Service) GetAllCoursesProjects() (*pgx.Rows, error) {
	return s.db.GetAllCoursesProjects()
}
