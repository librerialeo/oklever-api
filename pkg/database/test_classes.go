package database

import (
	"github.com/jackc/pgtype"

	"github.com/jackc/pgx"
)

// DBTestClass struct of database test_classes
type DBTestClass struct {
	ID     pgtype.Int4    `json:"id"`
	Name   pgtype.Varchar `json:"name"`
	Video  pgtype.Varchar `json:"video"`
	Status pgtype.Int4    `json:"status"`
}

// GetAllTestClasses queries for all testClasses
func (db *Database) GetAllTestClasses() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM test_classes")
}

// GetTestClassByTeachersID get test class by teacher id
func (db *Database) GetTestClassByTeachersID(teacherID int32) (*pgx.Rows, error) {
	return db.conn.Query("SELECT test_class_id, test_class_name, test_class_data, test_class_status FROM test_classes WHERE teacher_id = $1", teacherID)
}

// AddTeachersTestClass add new test class by teacher
func (db *Database) AddTeachersTestClass(teacherID int32, name string, video string, status int32) (*pgx.Rows, error) {
	return db.conn.Query("INSERT INTO test_classes(teacher_id,test_class_name, test_class_data, test_class_status) VALUES($1,$2,$3,$4) RETURNING test_class_id,test_class_name, test_class_data, test_class_status", teacherID, name, video, status)
}

// UpdateTeachersTestClass update teachers test class by teacher id
func (db *Database) UpdateTeachersTestClass(teacherID int32, name string, video string, status int32) (*pgx.Rows, error) {
	return db.conn.Query("UPDATE test_classes SET test_class_name = $2, test_class_data = $3, test_class_status = $4 WHERE teacher_id = $1 RETURNING test_class_id,test_class_name, test_class_data, test_class_status", teacherID, name, video, status)
}
