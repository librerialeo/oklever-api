package service

import (
	"github.com/jackc/pgx"
)

// GetAllUsersInsignias return all usersInsignias
func (s *Service) GetAllUsersInsignias() (*pgx.Rows, error) {
	return s.db.GetAllUsersInsignias()
}
