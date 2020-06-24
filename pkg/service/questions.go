package service

import (
	"github.com/jackc/pgx"
)

// Question struct question
type Question struct {
	ID       int32      `json:"id"`
	Type     int32      `json:"type"`
	Question string     `json:"question"`
	Options  []*Options `json:"options"`
}

// GetAllQuestions return all questions
func (s *Service) GetAllQuestions() (*pgx.Rows, error) {
	return s.db.GetAllQuestions()
}
