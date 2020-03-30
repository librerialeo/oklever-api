package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllTeachersExpertises() (pgx.Rows, error) {
	return s.db.getAllTeachersExpertises()
}