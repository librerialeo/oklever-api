package service

import (
	"github.com/jackc/pgx"
)

// GetAllCoursesReviews return all coursesReviews
func (s *Service) GetAllCoursesReviews() (*pgx.Rows, error) {
	return s.db.GetAllCoursesReviews()
}
