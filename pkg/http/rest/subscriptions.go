package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitSubscriptionsHandler initialize Subscriptions router
func InitSubscriptionsHandler(r *router.Router) {
	r.GET("/", getAllSubscriptions)
}

func getAllSubscriptions(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all Subscriptions")
}
