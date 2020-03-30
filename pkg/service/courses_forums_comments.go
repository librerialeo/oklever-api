package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllCoursesForumsComments() (pgx.Rows, error) {
	return s.db.getAllCoursesForumsComments()
}