package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitTestClassesHandler initialize TestClasses router
func InitTestClassesHandler(r *router.Router) {
	r.GET("/", getAllTestClasses)
}

func getAllTestClasses(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all TestClasses")
}
