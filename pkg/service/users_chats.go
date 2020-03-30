package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllUsersChats() (pgx.Rows, error) {
	return s.db.getAllUsersChats()
}