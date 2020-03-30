package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitSynchronousClassesResourcesHandler initialize SynchronousClassesResources router
func InitSynchronousClassesResourcesHandler(r *router.Router) {
	r.GET("/", getAllSynchronousClassesResources)
}

func getAllSynchronousClassesResources(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all SynchronousClassesResources")
}
