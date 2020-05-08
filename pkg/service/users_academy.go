package service

import (
	"github.com/jackc/pgx"
)

// GetAllUsersAcademy return all usersAcademy
func (s *Service) GetAllUsersAcademy(userID int32) (pgx.Rows, error) {
	return s.db.GetAllUsersAcademy(userID)
}
