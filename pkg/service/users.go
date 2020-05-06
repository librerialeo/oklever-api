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
func (s *Service) GetUserLastAction(userID int32) (time.Time, error) {
	rows, err := s.db.GetUserLastAction(userID)
	if err != nil {
		return time.Time{}, err
	}
	var lastaction pgtype.Timestamp
	if rows.Next() {
		err = rows.Scan(&lastaction)
	}
	if err != nil {
		return time.Time{}, err
	}
	return lastaction.Time, nil
}

// UpdateUserLastAction gets the last action time of the user with user id equal to userID
func (s *Service) UpdateUserLastAction(userID int32, lastaction time.Time) error {
	_, err := s.db.UpdateUserLastAction(userID, lastaction)
	return err
}
