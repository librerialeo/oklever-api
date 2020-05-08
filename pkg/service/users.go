package service

import (
	"time"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx"
)

// GetAllUsers return all users
func (s *Service) GetAllUsers() (pgx.Rows, error) {
	return s.db.GetAllUsers()
}

// GetUserLastAction gets the last action time of the user with user id equal to userID
func (s *Service) GetUserLastAction(userID int32) (*time.Time, error) {
	rows, err := s.db.GetUserLastAction(userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var lastaction pgtype.Timestamptz
	if rows.Next() {
		err = rows.Scan(&lastaction)
	}
	if err != nil {
		return nil, err
	}
	return &lastaction.Time, nil
}

// UpdateUserLastAction gets the last action time of the user with user id equal to userID
func (s *Service) UpdateUserLastAction(userID int32, lastaction time.Time) (*time.Time, error) {
	rows, err := s.db.UpdateUserLastAction(userID, lastaction)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var last pgtype.Timestamptz
	if rows.Next() {
		err = rows.Scan(&last)
		if err != nil {
			return nil, err
		}
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return &last.Time, nil
}

// UpdateUserInformation update user information
func (s *Service) UpdateUserInformation(userID int32, first string, last string, email string, gender string, phone string) error {
	rows, err := s.db.UpdateUserInformation(userID, first, last, email, gender, phone)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
