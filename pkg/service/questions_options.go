package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllQuestionsOptions() (pgx.Rows, error) {
	return s.db.getAllQuestionsOptions()
}