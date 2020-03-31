package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitTopicsResourcesHandler initialize topicsResources router
func InitTopicsResourcesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllTopicsResources(s))
}

func getAllTopicsResources(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all topicsResources")
		return err
	}
}
