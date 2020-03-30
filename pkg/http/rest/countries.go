package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitCountriesHandler initialize countries router
func InitCountriesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllCountries(s))
}

func getAllCountries(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all countries")
	}
}
