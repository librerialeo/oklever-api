package service

import (
	"github.com/librerialeo/oklever-api/pkg/database"
)

// AddAcademy asdf
func (s *Service) AddAcademy(firstname string, lastname string, email string, password string) (*database.DBUser, error) {
	rows, err := s.db.AddUser(firstname, lastname, email, password, 3)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var u database.DBUser
	if rows.Next() {
		err = rows.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email, &u.Rol)
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// GetAcademyUserByEmail asdf
func (s *Service) GetAcademyUserByEmail(email string) (*database.DBUser, error) {
	rows, err := s.db.GetAcademyByEmail(email)
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

// UpdateAcademyInformation asdf
func (s *Service) UpdateAcademyInformation(userID int32, license string, rfc string) error {
	return s.UpdateTeacherInformation(userID, license, rfc)
}
