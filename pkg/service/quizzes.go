package service

import (
	"github.com/jackc/pgx"
)

// Quiz struct return quiz
type Quiz struct {
	ID        int32       `json:"id"`
	Type      int32       `json:"type"`
	Name      string      `json:"name"`
	Questions []*Question `json:"questions"`
}

// GetAllQuizzes return all quizzes
func (s *Service) GetAllQuizzes() (*pgx.Rows, error) {
	return s.db.GetAllQuizzes()
}

// GetQuizByID return quiz by ID
func (s *Service) GetQuizByID(ID int32) (*Quiz, error) {
	rows, err := s.db.GetQuizByID(ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var qz Quiz
	var temp = make(map[int32]*Question)
	for rows.Next() {
		var q Question
		var op Options
		err = rows.Scan(&qz.ID, &qz.Type, &qz.Name, &q.ID, &q.Type, &q.Question, &op.ID, &op.Option)
		if err != nil {
			return nil, err
		}
		if temp[q.ID] == nil {
			temp[q.ID] = &q
			temp[q.ID].Options = append(temp[q.ID].Options, &op)
		} else {
			temp[q.ID].Options = append(temp[q.ID].Options, &op)
		}
	}
	for _, question := range temp {
		qz.Questions = append(qz.Questions, question)
	}
	return &qz, err
}
