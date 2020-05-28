package service

import (
	"github.com/librerialeo/oklever-api/pkg/database"
)

// GetUserAcademiesByUserID return all usersAcademy
func (s *Service) GetUserAcademiesByUserID(userID int32) (*[]database.DBUsersAcademy, error) {
	rows, err := s.db.GetUserAcademiesByUserID(userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	academies := []database.DBUsersAcademy{}
	err = nil
	for rows.Next() {
		var a database.DBUsersAcademy
		err = rows.Scan(
			&a.ID,
			&a.DegreeID,
			&a.UserAcademyName,
			&a.UserAcademyInstitution,
			&a.UserAcademyYear,
		)
		if err != nil {
			return nil, err
		}
		academies = append(academies, a)
	}
	return &academies, nil
}

// AddUsersAcademy add new academy user
func (s *Service) AddUsersAcademy(userID int32, degreeID int, academyName string, institution string, year int) (*[]database.DBUsersAcademy, error) {
	rows, err := s.db.AddUsersAcademy(userID, degreeID, academyName, institution, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	academies := []database.DBUsersAcademy{}
	err = nil
	if rows.Next() {
		var a database.DBUsersAcademy
		err = rows.Scan(&a.ID, &a.DegreeID, &a.UserAcademyName, &a.UserAcademyInstitution, &a.UserAcademyYear)
		if err != nil {
			return nil, err
		}
		academies = append(academies, a)
	}
	return &academies, nil
}

// UpdateUsersAcademy update users academy
func (s *Service) UpdateUsersAcademy(ID int, userID int32, degreeID int, academyName string, institution string, year int) (*[]database.DBUsersAcademy, error) {
	rows, err := s.db.UpdateUsersAcademy(ID, userID, degreeID, academyName, institution, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	academies := []database.DBUsersAcademy{}
	err = nil
	if rows.Next() {
		var a database.DBUsersAcademy
		err = rows.Scan(&a.ID, &a.DegreeID, &a.UserAcademyName, &a.UserAcademyInstitution, &a.UserAcademyYear)
		if err != nil {
			return nil, err
		}
		academies = append(academies, a)
	}
	return &academies, nil
}

// DeleteUsersAcademy  delete users academy by id
func (s *Service) DeleteUsersAcademy(ID int, userID int32) error {
	rows, err := s.db.DeleteUsersAcademy(ID, userID)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
