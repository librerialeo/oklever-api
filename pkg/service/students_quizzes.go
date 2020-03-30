package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllStudentsQuizzes() (pgx.Rows, error) {
	return s.db.getAllStudentsQuizzes()
}