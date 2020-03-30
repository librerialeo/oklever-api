package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllStudentsCourses() (pgx.Rows, error) {
	return s.db.getAllStudentsCourses()
}