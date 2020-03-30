package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitSynchronousClassesResourcesHandler initialize synchronousClassesResources router
func InitSynchronousClassesResourcesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllSynchronousClassesResources(s))
}

func getAllSynchronousClassesResources(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all synchronousClassesResources")
	}
}
