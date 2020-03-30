package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitClassesHandler initialize classes router
func InitClassesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllClasses(s))
}

func getAllClasses(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all classes")
	}
}
