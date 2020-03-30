package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllUsersDegrees() (pgx.Rows, error) {
	return s.db.getAllUsersDegrees()
}