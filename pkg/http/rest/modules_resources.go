package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitModulesResourcesHandler initialize modulesResources router
func InitModulesResourcesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllModulesResources(s))
}

func getAllModulesResources(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all modulesResources")
	}
}
