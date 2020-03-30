package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllTopics() (pgx.Rows, error) {
	return s.db.getAllTopics()
}