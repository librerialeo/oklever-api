package service

import (
	"github.com/jackc/pgx"
	"github.com/librerialeo/oklever-api/pkg/database"
)

func readManagement(rows *pgx.Rows) (*database.DBManagement, error) {
	if (*rows).Next() {
		var management database.DBManagement
		err := (*rows).Scan(&management.ID,
			&management.UserID,
			&management.Job,
			&management.Institution,
			&management.Months,
			&management.Added,
			&management.Modified,
			&management.Deleted)
		if err != nil {
			return nil, err
		}
		return &management, nil
	}
	return nil, nil
}

// GetUserManagement get the management of the passed id
func (s *Service) GetUserManagement(managementID int32) (*database.DBManagement, error) {
	rows, err := s.db.GetUserManagement(managementID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return readManagement(&rows)
}

// GetUserManagements get all user managements
func (s *Service) GetUserManagements(userID int32) (*[]database.DBManagement, error) {
	rows, err := s.db.GetUserManagements(userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	managements := []database.DBManagement{}
	management, err := readManagement(&rows)
	for management != nil && err == nil {
		managements = append(managements, *management)
		management, err = readManagement(&rows)
	}
	if err != nil {
		return nil, err
	}
	return &managements, nil
}

// AddUserManagement add a new user management
func (s *Service) AddUserManagement(userID int32, job string, institution string, months int16) (*database.DBManagement, error) {
	rows, err := s.db.AddUserManagement(userID, job, institution, months)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return readManagement(&rows)
}

// UpdateUserManagement update user management by id
func (s *Service) UpdateUserManagement(managementID int32, job string, institution string, months int16) (*database.DBManagement, error) {
	rows, err := s.db.UpdateUserManagement(managementID, job, institution, months)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return readManagement(&rows)
}

// DeleteUserManagement delete user management by id
func (s *Service) DeleteUserManagement(managementID int32) error {
	rows, err := s.db.DeleteUserManagement(managementID)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
