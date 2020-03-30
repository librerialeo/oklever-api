package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllTeachersManagements() (pgx.Rows, error) {
	return s.db.getAllTeachersManagements()
}