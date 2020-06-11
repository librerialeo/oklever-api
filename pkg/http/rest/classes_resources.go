package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo/v11"
)

// InitClassesResourcesHandler initialize ClassesResources router
func InitClassesResourcesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllClassesResources(s))
}

func getAllClassesResources(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all ClassesResources")
		return err
	}
}
