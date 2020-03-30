package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitModulesHandler initialize modules router
func InitModulesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllModules(s))
}

func getAllModules(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all modules")
	}
}
