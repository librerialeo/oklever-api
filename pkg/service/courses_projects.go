package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllCoursesProjects() (pgx.Rows, error) {
	return s.db.getAllCoursesProjects()
}