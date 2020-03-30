package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitLanguagesHandler initialize languages router
func InitLanguagesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllLanguages(s))
}

func getAllLanguages(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all languages")
	}
}
