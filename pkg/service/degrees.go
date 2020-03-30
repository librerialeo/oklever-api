package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllDegrees() (pgx.Rows, error) {
	return s.db.getAllDegrees()
}