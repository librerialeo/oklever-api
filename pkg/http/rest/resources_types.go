package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitResourcesTypesHandler initialize resourcesTypes router
func InitResourcesTypesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllResourcesTypes(s))
}

func getAllResourcesTypes(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all resourcesTypes")
		return err
	}
}
