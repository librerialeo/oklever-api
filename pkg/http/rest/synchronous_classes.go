package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitSynchronousClassesHandler initialize synchronousClasses router
func InitSynchronousClassesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllSynchronousClasses(s))
}

func getAllSynchronousClasses(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all synchronousClasses")
	}
}
