package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitTestClassesFeedbackHandler initialize testClassesFeedback router
func InitTestClassesFeedbackHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllTestClassesFeedback(s))
}

func getAllTestClassesFeedback(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all testClassesFeedback")
		return err
	}
}
