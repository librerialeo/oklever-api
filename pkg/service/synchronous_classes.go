package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllSynchronousClasses() (pgx.Rows, error) {
	return s.db.getAllSynchronousClasses()
}