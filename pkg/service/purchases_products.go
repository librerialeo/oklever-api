package service

import (
	"github.com/jackc/pgx"
)

// GetAllPurchasesProducts return all purchasesProducts
func (s *Service) GetAllPurchasesProducts() (*pgx.Rows, error) {
	return s.db.GetAllPurchasesProducts()
}
