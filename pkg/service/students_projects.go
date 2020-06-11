package service

import (
	"github.com/jackc/pgx"
)

// GetAllStudentsProjects return all studentsProjects
func (s *Service) GetAllStudentsProjects() (*pgx.Rows, error) {
	return s.db.GetAllStudentsProjects()
}
