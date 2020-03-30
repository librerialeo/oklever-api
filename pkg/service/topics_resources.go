package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllTopicsResources() (pgx.Rows, error) {
	return s.db.getAllTopicsResources()
}