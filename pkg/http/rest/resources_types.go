package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitResourcesTypesHandler initialize ResourcesTypes router
func InitResourcesTypesHandler(r *router.Router) {
	r.GET("/", getAllResourcesTypes)
}

func getAllResourcesTypes(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all ResourcesTypes")
}
