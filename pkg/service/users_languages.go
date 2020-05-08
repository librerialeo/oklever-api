package service

import (
	"github.com/jackc/pgx"
)

// GetAllUsersLanguages return all usersLanguages
func (s *Service) GetAllUsersLanguages(userID int32) (pgx.Rows, error) {
	return s.db.GetAllUsersLanguages(userID)
}

// AddUsersLanguages return success
func (s *Service) AddUsersLanguages(usersID int32, languageID int) error {
	rows, err := s.db.AddUsersLanguages(usersID, languageID)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

// DeleteUsersLanguages delete user language
func (s *Service) DeleteUsersLanguages(userID int32, languageID int) error {
	rows, err := s.db.DeleteUsersLanguages(userID, languageID)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
