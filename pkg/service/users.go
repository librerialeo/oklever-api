package service

import (
	"github.com/jackc/pgx"
)

// GetAllUsers return all users
func (s *Service) GetAllUsers() (pgx.Rows, error) {
	return s.db.GetAllUsers()
}
