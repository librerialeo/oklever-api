package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitStudentsQuizzesHandler initialize studentsQuizzes router
func InitStudentsQuizzesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllStudentsQuizzes(s))
}

func getAllStudentsQuizzes(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all studentsQuizzes")
		return err
	}
}
