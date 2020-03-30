package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitClassesResourcesHandler initialize ClassesResources router
func InitClassesResourcesHandler(r *router.Router) {
	r.GET("/", getAllClassesResources)
}

func getAllClassesResources(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all ClassesResources")
}
