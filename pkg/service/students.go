package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllStudents() (pgx.Rows, error) {
	return s.db.getAllStudents()
}