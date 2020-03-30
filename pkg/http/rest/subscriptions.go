package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitSubscriptionsHandler initialize subscriptions router
func InitSubscriptionsHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllSubscriptions(s))
}

func getAllSubscriptions(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all subscriptions")
	}
}
