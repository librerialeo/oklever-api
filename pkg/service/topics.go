package service

import (
	"github.com/jackc/pgx"
)

// GetAllTopics return all topics
func (s *Service) GetAllTopics() (pgx.Rows, error) {
	return s.db.GetAllTopics()
}
