package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitModulesFeedbackHandler initialize modulesFeedback router
func InitModulesFeedbackHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllModulesFeedback(s))
}

func getAllModulesFeedback(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all modulesFeedback")
		return err
	}
}
