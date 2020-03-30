package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitCountriesHandler initialize Countries router
func InitCountriesHandler(r *router.Router) {
	r.GET("/", getAllCountries)
}

func getAllCountries(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all Countries")
}
