package service

import (
	"github.com/jackc/pgx"
)

// GetAllStudentsProjectsHistory return all studentsProjectsHistory
func (s *Service) GetAllStudentsProjectsHistory() (pgx.Rows, error) {
	return s.db.GetAllStudentsProjectsHistory()
}
