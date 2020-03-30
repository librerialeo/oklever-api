package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllUsers() (pgx.Rows, error) {
	return s.db.getAllUsers()
}