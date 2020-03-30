package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllClassesComments() (pgx.Rows, error) {
	return s.db.getAllClassesComments()
}