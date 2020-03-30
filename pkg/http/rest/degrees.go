package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitDegreesHandler initialize degrees router
func InitDegreesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllDegrees(s))
}

func getAllDegrees(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all degrees")
	}
}
