package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllRoles() (pgx.Rows, error) {
	return s.db.getAllRoles()
}