package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitModulesHandler initialize modules router
func InitModulesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllModules(s))
}

func getAllModules(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all modules")
		return err
	}
}
