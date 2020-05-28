package service

import (
	"github.com/jackc/pgx"
	"github.com/librerialeo/oklever-api/pkg/database"
)

// GetAllTestClasses return all testClasses
func (s *Service) GetAllTestClasses() (pgx.Rows, error) {
	return s.db.GetAllTestClasses()
}

// GetTestClassByTeacherID get test class by teacher id
func (s *Service) GetTestClassByTeacherID(teacherID int32) (*database.DBTestClass, error) {
	rows, err := s.db.GetTestClassByTeachersID(teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var t database.DBTestClass
	if rows.Next() {
		err = rows.Scan(t.Name, t.Video)
	}
	return &t, err
}

// AddTeachersTestClass add test class for teachers
func (s *Service) AddTeachersTestClass(teacherID int32, name string, video string) (*database.DBTestClass, error) {
	rows, err := s.db.AddTeachersTestClass(teacherID, name, video)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var t database.DBTestClass
	if rows.Next() {
		err = rows.Scan(t.Name, t.Video)
	}
	return &t, err
}

// UpdateTeachersTestClass update teachers test class
func (s *Service) UpdateTeachersTestClass(teacherID int32, name string, video string) (*database.DBTestClass, error) {
	rows, err := s.db.UpdateTeachersTestClass(teacherID, name, video)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var t database.DBTestClass
	if rows.Next() {
		err = rows.Scan(t.Name, t.Video)
	}
	return &t, err
}
