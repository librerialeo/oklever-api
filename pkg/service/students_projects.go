package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllStudentsProjects() (pgx.Rows, error) {
	return s.db.getAllStudentsProjects()
}