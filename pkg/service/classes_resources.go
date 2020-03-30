package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllClassesResources() (pgx.Rows, error) {
	return s.db.getAllClassesResources()
}