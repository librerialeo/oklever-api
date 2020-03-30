package service

import (
	"github.com/jackc/pgx"
	"github.com/librerialeo/oklever-api/pkg/database"
)

// Service is a wrapper for models actions
type Service struct {
	db *database.Database
}

// InitService Initialize models actions
func InitService(conn *pgx.Conn) *Service {
	return &Service{database.InitDatabase(conn)}
}
