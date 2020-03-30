package service

import (
	"github.com/jackc/pgx"
)

// GetAllLanguages return all languages
func (s *Service) GetAllLanguages() (pgx.Rows, error) {
	return s.db.GetAllLanguages()
}
