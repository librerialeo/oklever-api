package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitDegreesHandler initialize Degrees router
func InitDegreesHandler(r *router.Router) {
	r.GET("/", getAllDegrees)
}

func getAllDegrees(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all Degrees")
}
