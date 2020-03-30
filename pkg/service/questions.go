package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllQuestions() (pgx.Rows, error) {
	return s.db.getAllQuestions()
}