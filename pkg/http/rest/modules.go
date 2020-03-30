package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitModulesHandler initialize Modules router
func InitModulesHandler(r *router.Router) {
	r.GET("/", getAllModules)
}

func getAllModules(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all Modules")
}
