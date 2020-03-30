package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllTestClasses() (pgx.Rows, error) {
	return s.db.getAllTestClasses()
}