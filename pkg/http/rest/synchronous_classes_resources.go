package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo/v11"
)

// InitSynchronousClassesResourcesHandler initialize synchronousClassesResources router
func InitSynchronousClassesResourcesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllSynchronousClassesResources(s))
}

func getAllSynchronousClassesResources(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all synchronousClassesResources")
		return err
	}
}
