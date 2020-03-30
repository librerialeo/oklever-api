package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitModulesResourcesHandler initialize ModulesResources router
func InitModulesResourcesHandler(r *router.Router) {
	r.GET("/", getAllModulesResources)
}

func getAllModulesResources(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all ModulesResources")
}
