package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo/v11"
)

// InitQuestionsOptionsHandler initialize questionsOptions router
func InitQuestionsOptionsHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllQuestionsOptions(s))
}

func getAllQuestionsOptions(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all questionsOptions")
		return err
	}
}
