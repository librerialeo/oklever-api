package service

import (
	"time"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx"
	"github.com/librerialeo/oklever-api/pkg/database"
)

// UserCredential user credential info
type UserCredential struct {
	ID     int32
	Rol    int32
	Status string
}

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

// GetUserBiography get user biography string
func (s *Service) GetUserBiography(userID int32) (string, error) {
	rows, err := s.db.GetUserBiography(userID)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var biography pgtype.Text
	if rows.Next() {
		err := rows.Scan(&biography)
		if err != nil {
			return "", nil
		}
	}
	return biography.String, nil
}

// SetUserBiography get user biography string
func (s *Service) SetUserBiography(userID int32, biography string) (string, error) {
	rows, err := s.db.SetUserBiography(userID, biography)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var DBbiography pgtype.Text
	if rows.Next() {
		err := rows.Scan(&DBbiography)
		if err != nil {
			return "", nil
		}
	}
	return DBbiography.String, nil
}

// SetUserImage get user biography string
func (s *Service) SetUserImage(userID int32, image string) (string, error) {
	rows, err := s.db.SetUserImage(userID, image)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var DBimage pgtype.Varchar
	if rows.Next() {
		err := rows.Scan(&DBimage)
		if err != nil {
			return "", nil
		}
	}
	return DBimage.String, nil
}

// SetUserStatus set user status
func (s *Service) SetUserStatus(userID int32, status string) error {
	rows, err := s.db.SetUserStatus(userID, status)
	if err != nil {
		return err
	}
	rows.Close()
	return nil
}

// GetUserByID get user with passed id
func (s *Service) GetUserByID(userID int32) (*database.DBUser, error) {
	rows, err := s.db.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var u database.DBUser
	if rows.Next() {
		err = rows.Scan(&u.ID,
			&u.Email,
			&u.Password,
			&u.Firstname,
			&u.Lastname,
			&u.Gender,
			&u.Image,
			&u.Birthdate,
			&u.Phone,
			&u.License,
			&u.RFC,
			&u.Biography,
			&u.TeachingMonths,
			&u.Status,
			&u.Country,
			&u.Rol,
			&u.LastAction,
			&u.Created,
			&u.Modified,
			&u.Deleted)
	}
	return &u, err
}

// GetUserCredentialsByID get user credentials data
func (s *Service) GetUserCredentialsByID(userID int32) (*UserCredential, error) {
	rows, err := s.db.GetUserCredentialsByID(userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var u database.DBUser
	var credential UserCredential
	if rows.Next() {
		err = rows.Scan(&u.ID,
			&u.Rol,
			&u.Status)
		if err != nil {
			return nil, err
		}
		credential = UserCredential{ID: u.ID.Int, Rol: u.Rol.Int, Status: u.Status.String}
	}
	return &credential, err
}
