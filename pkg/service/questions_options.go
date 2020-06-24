package service

import (
	"github.com/jackc/pgx"
)

// Options return struct database of options for question
type Options struct {
	ID     int32  `json:"id"`
	Option string `json:"option"`
}

// GetAllQuestionsOptions return all questionsOptions
func (s *Service) GetAllQuestionsOptions() (*pgx.Rows, error) {
	return s.db.GetAllQuestionsOptions()
}
