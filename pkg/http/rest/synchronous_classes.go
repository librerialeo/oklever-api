package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitSynchronousClassesHandler initialize synchronousClasses router
func InitSynchronousClassesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllSynchronousClasses(s))
}

func getAllSynchronousClasses(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all synchronousClasses")
		return err
	}
}
