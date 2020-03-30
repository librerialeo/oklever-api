package service

import (
	"github.com/jackc/pgx"
)

// GetAllRoles return all roles
func (s *Service) GetAllRoles() (pgx.Rows, error) {
	return s.db.GetAllRoles()
}
