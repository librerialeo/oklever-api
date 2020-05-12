package service

import (
	"github.com/librerialeo/oklever-api/pkg/database"
)

// GetAllDegrees return all degrees
func (s *Service) GetAllDegrees() (*[]database.DBDegrees, error) {
	var err error
	rows, err := s.db.GetAllDegrees()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	degrees := make([]database.DBDegrees, 0)
	err = nil
	for rows.Next() {
		var d database.DBDegrees
		err = rows.Scan(&d.ID, &d.Name, &d.DegreeClass)
		if err != nil {
			return nil, err
		}
		degrees = append(degrees, d)
	}
	return &degrees, nil
}
