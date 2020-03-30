package service

import (
	"github.com/jackc/pgx"
)

// GetAllUsersChats return all usersChats
func (s *Service) GetAllUsersChats() (pgx.Rows, error) {
	return s.db.GetAllUsersChats()
}
