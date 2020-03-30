package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllStudentsAnswers() (pgx.Rows, error) {
	return s.db.getAllStudentsAnswers()
}