package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllTeachersResearch() (pgx.Rows, error) {
	return s.db.getAllTeachersResearch()
}