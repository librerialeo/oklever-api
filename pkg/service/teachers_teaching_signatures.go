package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllTeachersTeachingSignatures() (pgx.Rows, error) {
	return s.db.getAllTeachersTeachingSignatures()
}