package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo/v11"
)

// InitRolesHandler initialize roles router
func InitRolesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllRoles(s))
}

func getAllRoles(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all roles")
		return err
	}
}
