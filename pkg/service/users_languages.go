package service

import (
	"github.com/jackc/pgtype"
)

// GetAllUsersLanguages return all usersLanguages
func (s *Service) GetAllUsersLanguages(userID int32) (*[]int32, error) {
	rows, err := s.db.GetAllUsersLanguages(userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	languages := []int32{}
	for rows.Next() {
		var l pgtype.Int4
		err = rows.Scan(&l)
		if err != nil {
			return nil, err
		}
		languages = append(languages, l.Int)
	}
	return &languages, nil
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
