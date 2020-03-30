package service

import (
	"github.com/jackc/pgx"
)

// GetAllCoursesDatasheets return all coursesDatasheets
func (s *Service) GetAllCoursesDatasheets() (pgx.Rows, error) {
	return s.db.GetAllCoursesDatasheets()
}
