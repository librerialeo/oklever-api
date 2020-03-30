package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllTeachersTeaching() (pgx.Rows, error) {
	return s.db.getAllTeachersTeaching()
}