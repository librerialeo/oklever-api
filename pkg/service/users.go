package service

import (
	"fmt"
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
	fmt.Println("get", lastaction)
	return &lastaction.Time, nil
}

// UpdateUserLastAction gets the last action time of the user with user id equal to userID
func (s *Service) UpdateUserLastAction(userID int32, lastaction time.Time) (*time.Time, error) {
	fmt.Println("set", userID, lastaction)
	rows, err := s.db.UpdateUserLastAction(userID, lastaction)
	fmt.Println("rows:", rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var last *pgtype.Timestamptz
	if rows.Next() {
		err = rows.Scan(last)
		return nil, err
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return &last.Time, nil
}
