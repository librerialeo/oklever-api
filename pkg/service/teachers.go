package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllTeachers() (pgx.Rows, error) {
	return s.db.getAllTeachers()
}