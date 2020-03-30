package service

import (
	"github.com/jackc/pgx"
)

func (s *Service) getAllPurchases() (pgx.Rows, error) {
	return s.db.getAllPurchases()
}