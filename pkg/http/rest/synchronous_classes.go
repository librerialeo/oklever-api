package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitSynchronousClassesHandler initialize SynchronousClasses router
func InitSynchronousClassesHandler(r *router.Router) {
	r.GET("/", getAllSynchronousClasses)
}

func getAllSynchronousClasses(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all SynchronousClasses")
}
