package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitResourcesTypesHandler initialize resourcesTypes router
func InitResourcesTypesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllResourcesTypes(s))
}

func getAllResourcesTypes(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all resourcesTypes")
	}
}
