package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitStudentsAnswersHandler initialize studentsAnswers router
func InitStudentsAnswersHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllStudentsAnswers(s))
}

func getAllStudentsAnswers(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all studentsAnswers")
		return err
	}
}
