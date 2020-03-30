package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitInsigniasHandler initialize insignias router
func InitInsigniasHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllInsignias(s))
}

func getAllInsignias(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all insignias")
	}
}
