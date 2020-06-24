package database

import (
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx"
)

// DBQuiz struct of database quizzes
type DBQuiz struct {
	ID           pgtype.Int4    `json:"id"`
	QuizType     pgtype.Int4    `json:"quiz_type"`
	QuizName     pgtype.Varchar `json:"name"`
	QuestionID   pgtype.Int4    `json:"question_id"`
	QuestionType pgtype.Int4    `json:"question_type"`
	Question     pgtype.Varchar `json:"question"`
	Option       pgtype.Varchar `json:"option"`
}

// GetAllQuizzes queries for all quizzes
func (db *Database) GetAllQuizzes() (*pgx.Rows, error) {
	return db.conn.Query("SELECT * FROM quizzes")
}

// GetQuizByID get quiz by id of quizzes
func (db *Database) GetQuizByID(ID int32) (*pgx.Rows, error) {
	return db.conn.Query("SELECT qz.quiz_id, qz.quiz_type_id, qzt.quiz_type_name, q.question_id, q.question_type_id, q.question, qo.question_option_id, qo.question_option FROM quizzes qz JOIN quizzes_types qzt ON (qz.quiz_type_id = qzt.quiz_type_id) JOIN questions q ON (q.quiz_id = qz.quiz_id) JOIN questions_options qo ON (q.question_id = qo.question_id) WHERE qz.quiz_id = $1", ID)
}
