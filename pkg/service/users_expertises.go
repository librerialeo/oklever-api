package service

import (
	"github.com/jackc/pgx"
	"github.com/librerialeo/oklever-api/pkg/database"
)

func readExpertise(rows *pgx.Rows) (*database.DBExpertise, error) {
	if (*rows).Next() {
		var expertise database.DBExpertise
		err := (*rows).Scan(&expertise.ID,
			&expertise.UserID,
			&expertise.Name,
			&expertise.Months,
			&expertise.Added,
			&expertise.Modified,
			&expertise.Deleted)
		if err != nil {
			return nil, err
		}
		return &expertise, nil
	}
	return nil, nil
}

// GetUserExpertise get the expertise of the passed id
func (s *Service) GetUserExpertise(expertiseID int32) (*database.DBExpertise, error) {
	rows, err := s.db.GetUserExpertise(expertiseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return readExpertise(&rows)
}

// GetUserExpertises get all user expertises
func (s *Service) GetUserExpertises(userID int32) (*[]database.DBExpertise, error) {
	rows, err := s.db.GetUserExpertises(userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	expertises := []database.DBExpertise{}
	expertise, err := readExpertise(&rows)
	for expertise != nil && err == nil {
		expertises = append(expertises, *expertise)
		expertise, err = readExpertise(&rows)
	}
	if err != nil {
		return nil, err
	}
	return &expertises, nil
}

// AddUserExpertise add a new user expertise
func (s *Service) AddUserExpertise(userID int32, name string, months int32) (*database.DBExpertise, error) {
	rows, err := s.db.AddUserExpertise(userID, name, months)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return readExpertise(&rows)
}

// UpdateUserExpertise update user expertise by id
func (s *Service) UpdateUserExpertise(expertiseID int32, name string, months int32) (*database.DBExpertise, error) {
	rows, err := s.db.UpdateUserExpertise(expertiseID, name, months)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return readExpertise(&rows)
}

// DeleteUserExpertise delete user expertise by id
func (s *Service) DeleteUserExpertise(expertiseID int32) error {
	rows, err := s.db.DeleteUserExpertise(expertiseID)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
