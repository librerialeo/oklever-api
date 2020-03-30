package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllCoursesForums() (pgx.Rows, error) {
	return s.db.getAllCoursesForums()
}