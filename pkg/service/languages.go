package service

import "github.com/librerialeo/oklever-api/pkg/database"

// GetAllLanguages return all languages
func (s *Service) GetAllLanguages() (*[]database.DBLanguage, error) {
	var err error
	rows, err := s.db.GetAllLanguages()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	languages := make([]database.DBLanguage, 0)
	err = nil
	for rows.Next() {
		var l database.DBLanguage
		err = rows.Scan(
			&l.ID,
			&l.Name,
			&l.ISO,
			&l.Created,
			&l.Modified,
			&l.Deleted)
		if err != nil {
			return nil, err
		}
		languages = append(languages, l)
	}
	return &languages, nil
}
