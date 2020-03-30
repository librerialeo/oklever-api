package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllQuizzes() (pgx.Rows, error) {
	return s.db.getAllQuizzes()
}