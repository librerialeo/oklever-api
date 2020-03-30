package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllQuestionsTypes() (pgx.Rows, error) {
	return s.db.getAllQuestionsTypes()
}