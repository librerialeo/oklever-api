package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllCoursesReviews() (pgx.Rows, error) {
	return s.db.getAllCoursesReviews()
}