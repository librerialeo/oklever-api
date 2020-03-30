package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitTestClassesHandler initialize testClasses router
func InitTestClassesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllTestClasses(s))
}

func getAllTestClasses(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all testClasses")
	}
}
