package service

import (
	"github.com/jackc/pgx"
)

// GetAllInsignias return all insignias
func (s *Service) GetAllInsignias() (pgx.Rows, error) {
	return s.db.GetAllInsignias()
}
