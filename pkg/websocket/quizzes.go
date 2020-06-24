package websocket

import (
	"fmt"
)

// GetAllQuizzes get all quizzes of table quizzes
func GetAllQuizzes(s *Socket, a *Action) {
	fmt.Println("inicio")
}

// GetQuizByID get quiz by id of quizzes
func GetQuizByID(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok {
		id, idOk := data["id"].(float64)
		if idOk {
			quiz, err := s.io.service.GetQuizByID(int32(id))
			if err != nil {
				s.EmitServerError("Error al obtener el quiz", err)
			} else {
				s.Emit("ADD_QUIZZES", quiz)
			}
		}
	}
}
