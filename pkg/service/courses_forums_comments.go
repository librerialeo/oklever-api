package service

import (
	"github.com/jackc/pgx"
)

// GetAllCoursesForumsComments return all coursesForumsComments
func (s *Service) GetAllCoursesForumsComments() (*pgx.Rows, error) {
	return s.db.GetAllCoursesForumsComments()
}
