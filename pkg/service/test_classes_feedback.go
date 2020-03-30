package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllTestClassesFeedback() (pgx.Rows, error) {
	return s.db.getAllTestClassesFeedback()
}