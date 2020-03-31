package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitQuizzesHandler initialize quizzes router
func InitQuizzesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllQuizzes(s))
}

func getAllQuizzes(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all quizzes")
		return err
	}
}
