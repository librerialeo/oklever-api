package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllTeachersLanguages() (pgx.Rows, error) {
	return s.db.getAllTeachersLanguages()
}