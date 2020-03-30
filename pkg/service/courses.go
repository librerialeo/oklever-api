package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllCourses() (pgx.Rows, error) {
	return s.db.getAllCourses()
}