package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitClassesHandler initialize Classes router
func InitClassesHandler(r *router.Router) {
	r.GET("/", getAllClasses)
}

func getAllClasses(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all Classes")
}
