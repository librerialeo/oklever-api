package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllCoursesDatasheets() (pgx.Rows, error) {
	return s.db.getAllCoursesDatasheets()
}