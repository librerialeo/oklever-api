package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllClasses() (pgx.Rows, error) {
	return s.db.getAllClasses()
}