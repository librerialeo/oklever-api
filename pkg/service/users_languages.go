package service

import (
	"github.com/jackc/pgx"
)

// GetAllUsersLanguages return all teachersLanguages
func (s *Service) GetAllUsersLanguages() (pgx.Rows, error) {
	return s.db.GetAllUsersLanguages()
}

// AddUsersLanguages return success
func (s *Service) AddUsersLanguages(teacherID int, languageID int) error {
	rows, err := s.db.AddUsersLanguages(teacherID, languageID)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

// DeleteUsersLanguages delete user language
func (s *Service) DeleteUsersLanguages(userID int, languageID int) error {
	rows, err := s.db.DeleteUsersLanguages(userID, languageID)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
