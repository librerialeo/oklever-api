package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitTestClassesHandler initialize testClasses router
func InitTestClassesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllTestClasses(s))
}

func getAllTestClasses(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all testClasses")
		return err
	}
}
