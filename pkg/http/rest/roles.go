package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitRolesHandler initialize roles router
func InitRolesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllRoles(s))
}

func getAllRoles(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all roles")
	}
}
