package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo/v11"
)

// InitQuestionsHandler initialize questions router
func InitQuestionsHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllQuestions(s))
}

func getAllQuestions(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all questions")
		return err
	}
}
