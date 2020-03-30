package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllChats() (pgx.Rows, error) {
	return s.db.getAllChats()
}