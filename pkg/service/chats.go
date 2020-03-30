package service

import (
	"github.com/jackc/pgx"
)

// GetAllChats return all chats
func (s *Service) GetAllChats() (pgx.Rows, error) {
	return s.db.GetAllChats()
}
