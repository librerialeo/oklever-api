package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllPurchasesProducts() (pgx.Rows, error) {
	return s.db.getAllPurchasesProducts()
}