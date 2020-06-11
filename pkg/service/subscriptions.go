package service

import (
	"github.com/jackc/pgx"
)

// GetAllSubscriptions return all subscriptions
func (s *Service) GetAllSubscriptions() (*pgx.Rows, error) {
	return s.db.GetAllSubscriptions()
}
