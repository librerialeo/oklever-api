package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitClassesResourcesHandler initialize ClassesResources router
func InitClassesResourcesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllClassesResources(s))
}

func getAllClassesResources(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all ClassesResources")
	}
}
