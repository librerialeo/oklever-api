package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitClassesHandler initialize classes router
func InitClassesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllClasses(s))
}

func getAllClasses(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all classes")
		return err
	}
}
