package service

import (
	"github.com/jackc/pgx"
	"github.com/librerialeo/oklever-api/pkg/database"
)

func readInvestment(rows *pgx.Rows) (*database.DBInvestment, error) {
	if (*rows).Next() {
		var investment database.DBInvestment
		err := (*rows).Scan(&investment.ID,
			&investment.UserID,
			&investment.Type,
			&investment.Reference,
			&investment.Year,
			&investment.Added,
			&investment.Modified,
			&investment.Deleted)
		if err != nil {
			return nil, err
		}
		return &investment, nil
	}
	return nil, nil
}

// GetUserInvestment get the investment of the passed id
func (s *Service) GetUserInvestment(investmentID int32) (*database.DBInvestment, error) {
	rows, err := s.db.GetUserInvestment(investmentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return readInvestment(&rows)
}

// GetUserInvestments get all user investments
func (s *Service) GetUserInvestments(userID int32) (*[]database.DBInvestment, error) {
	rows, err := s.db.GetUserInvestments(userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	investments := []database.DBInvestment{}
	investment, err := readInvestment(&rows)
	for investment != nil && err == nil {
		investments = append(investments, *investment)
		investment, err = readInvestment(&rows)
	}
	if err != nil {
		return nil, err
	}
	return &investments, nil
}

// AddUserInvestment add a new user investment
func (s *Service) AddUserInvestment(userID int32, typeof string, reference string, year int32) (*database.DBInvestment, error) {
	rows, err := s.db.AddUserInvestment(userID, typeof, reference, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return readInvestment(&rows)
}

// UpdateUserInvestment update user investment by id
func (s *Service) UpdateUserInvestment(investmentID int32, reference string, year int32) (*database.DBInvestment, error) {
	rows, err := s.db.UpdateUserInvestment(investmentID, reference, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return readInvestment(&rows)
}

// DeleteUserInvestment delete user investment by id
func (s *Service) DeleteUserInvestment(investmentID int32) error {
	rows, err := s.db.DeleteUserInvestment(investmentID)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
