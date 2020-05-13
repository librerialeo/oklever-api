package service

import (
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx"
	"github.com/librerialeo/oklever-api/pkg/database"
)

func readSignature(rows *pgx.Rows) (*database.DBSignature, error) {
	if (*rows).Next() {
		var signature database.DBSignature
		err := (*rows).Scan(&signature.ID,
			&signature.UserID,
			&signature.DegreeID,
			&signature.Name,
			&signature.Added,
			&signature.Modified,
			&signature.Deleted)
		if err != nil {
			return nil, err
		}
		return &signature, nil
	}
	return nil, nil
}

// GetUserTeachingSignature get the signature of the passed id
func (s *Service) GetUserTeachingSignature(signatureID int32) (*database.DBSignature, error) {
	rows, err := s.db.GetUserTeachingSignature(signatureID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return readSignature(&rows)
}

// GetUserTeachingSignatures get all user signatures
func (s *Service) GetUserTeachingSignatures(userID int32) (*[]database.DBSignature, error) {
	rows, err := s.db.GetUserTeachingSignatures(userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	signatures := []database.DBSignature{}
	signature, err := readSignature(&rows)
	for signature != nil && err == nil {
		signatures = append(signatures, *signature)
		signature, err = readSignature(&rows)
	}
	if err != nil {
		return nil, err
	}
	return &signatures, nil
}

// AddUserTeachingSignature add a new user signature
func (s *Service) AddUserTeachingSignature(userID int32, degreeID int32, name string) (*database.DBSignature, error) {
	rows, err := s.db.AddUserTeachingSignature(userID, degreeID, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return readSignature(&rows)
}

// UpdateUserTeachingSignature update user signature by id
func (s *Service) UpdateUserTeachingSignature(signatureID int32, name string) (*database.DBSignature, error) {
	rows, err := s.db.UpdateUserTeachingSignature(signatureID, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return readSignature(&rows)
}

// DeleteUserTeachingSignature delete user signature by id
func (s *Service) DeleteUserTeachingSignature(signatureID int32) error {
	rows, err := s.db.DeleteUserTeachingSignature(signatureID)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func readInstitution(rows *pgx.Rows) (*database.DBInstitution, error) {
	if (*rows).Next() {
		var institution database.DBInstitution
		err := (*rows).Scan(&institution.ID,
			&institution.UserID,
			&institution.Name,
			&institution.Added,
			&institution.Modified,
			&institution.Deleted)
		if err != nil {
			return nil, err
		}
		return &institution, nil
	}
	return nil, nil
}

// GetUserTeachingInstitution get the institution of the passed id
func (s *Service) GetUserTeachingInstitution(institutionID int32) (*database.DBInstitution, error) {
	rows, err := s.db.GetUserTeachingInstitution(institutionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return readInstitution(&rows)
}

// GetUserTeachingInstitutions get all user institutions
func (s *Service) GetUserTeachingInstitutions(userID int32) (*[]database.DBInstitution, error) {
	rows, err := s.db.GetUserTeachingInstitutions(userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	institutions := []database.DBInstitution{}
	institution, err := readInstitution(&rows)
	for institution != nil && err == nil {
		institutions = append(institutions, *institution)
		institution, err = readInstitution(&rows)
	}
	if err != nil {
		return nil, err
	}
	return &institutions, nil
}

// AddUserTeachingInstitution add a new user institution
func (s *Service) AddUserTeachingInstitution(userID int32, name string) (*database.DBInstitution, error) {
	rows, err := s.db.AddUserTeachingInstitution(userID, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return readInstitution(&rows)
}

// UpdateUserTeachingInstitution update user institution by id
func (s *Service) UpdateUserTeachingInstitution(institutionID int32, name string) (*database.DBInstitution, error) {
	rows, err := s.db.UpdateUserTeachingInstitution(institutionID, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return readInstitution(&rows)
}

// DeleteUserTeachingInstitution delete user institution by id
func (s *Service) DeleteUserTeachingInstitution(institutionID int32) error {
	rows, err := s.db.DeleteUserTeachingInstitution(institutionID)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

// GetUserExperience get the user experience of the passed id
func (s *Service) GetUserExperience(userID int32) (*int32, error) {
	rows, err := s.db.GetUserExperience(userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var experience pgtype.Int4
	if rows.Next() {
		err = rows.Scan(&experience)
		if err != nil {
			return nil, err
		}
	}
	return &experience.Int, nil
}

// SetUserExperience set the user experience of the passed id
func (s *Service) SetUserExperience(userID int32, months int32) (*int32, error) {
	rows, err := s.db.SetUserExperience(userID, months)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var experience pgtype.Int4
	if rows.Next() {
		err = rows.Scan(&experience)
		if err != nil {
			return nil, err
		}
	}
	return &experience.Int, nil
}
