package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitStudentsSubscriptionsHandler initialize StudentsSubscriptions router
func InitStudentsSubscriptionsHandler(r *router.Router) {
	r.GET("/", getAllStudentsSubscriptions)
}

func getAllStudentsSubscriptions(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all StudentsSubscriptions")
}
