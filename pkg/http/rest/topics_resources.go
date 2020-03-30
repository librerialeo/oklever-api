package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitTopicsResourcesHandler initialize topicsResources router
func InitTopicsResourcesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllTopicsResources(s))
}

func getAllTopicsResources(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all topicsResources")
	}
}
