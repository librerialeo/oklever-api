package service

import (
	"fmt"

	"github.com/librerialeo/oklever-api/pkg/database"
)

// GetAllUsersAcademy return all usersAcademy
func (s *Service) GetAllUsersAcademy(userID int32) (*[]database.DBUsersAcademy, error) {
	rows, err := s.db.GetAllUsersAcademy(userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	academies := make([]database.DBUsersAcademy, 0)
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
	fmt.Println(err, "err")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	academies := make([]database.DBUsersAcademy, 0)
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
