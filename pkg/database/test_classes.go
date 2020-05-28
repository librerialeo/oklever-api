package database

import (
	"context"

	"github.com/jackc/pgtype"

	"github.com/jackc/pgx"
)

// DBTestClass struct of database test_classes
type DBTestClass struct {
	Name  pgtype.Varchar `json:"name"`
	Video pgtype.Varchar `json:"video"`
}

// GetAllTestClasses queries for all testClasses
func (db *Database) GetAllTestClasses() (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT * FROM test_classes")
}

// GetTestClassByTeachersID get test class by teacher id
func (db *Database) GetTestClassByTeachersID(teacherID int32) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "SELECT test_class_name, test_class_data FROM test_classes WHERE teacher_id = $1", teacherID)
}

// AddTeachersTestClass add new test class by teacher
func (db *Database) AddTeachersTestClass(teacherID int32, name string, video string) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "INSERT INTO test_classes(teacher_id,test_class_name, test_class_data) VALUES($1,$2,$3) RETURNING test_class_name, test_class_data", teacherID, name, video)
}

// UpdateTeachersTestClass update teachers test class by teacher id
func (db *Database) UpdateTeachersTestClass(teacherID int32, name string, video string) (pgx.Rows, error) {
	return db.conn.Query(context.Background(), "UPDATE test_classes SET test_class_name = $2, test_class_data = $3 WHERE teacher_id = $1 RETURNING test_class_name, test_class_data", teacherID, name, video)
}
