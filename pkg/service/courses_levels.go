package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllCoursesLevels() (pgx.Rows, error) {
	return s.db.getAllCoursesLevels()
}