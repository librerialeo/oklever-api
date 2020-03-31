package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitModulesResourcesHandler initialize modulesResources router
func InitModulesResourcesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllModulesResources(s))
}

func getAllModulesResources(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all modulesResources")
		return err
	}
}
