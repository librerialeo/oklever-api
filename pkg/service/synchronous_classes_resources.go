package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllSynchronousClassesResources() (pgx.Rows, error) {
	return s.db.getAllSynchronousClassesResources()
}