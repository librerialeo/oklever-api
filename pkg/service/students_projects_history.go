package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllStudentsProjectsHistory() (pgx.Rows, error) {
	return s.db.getAllStudentsProjectsHistory()
}