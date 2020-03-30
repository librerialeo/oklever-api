package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllTeachersTeachingInstitutions() (pgx.Rows, error) {
	return s.db.getAllTeachersTeachingInstitutions()
}