package service

import (
	"github.com/jackc/pgx"
)

// GetAllPurchases return all purchases
func (s *Service) GetAllPurchases() (*pgx.Rows, error) {
	return s.db.GetAllPurchases()
}
