package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitStudentsSubscriptionsHandler initialize studentsSubscriptions router
func InitStudentsSubscriptionsHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllStudentsSubscriptions(s))
}

func getAllStudentsSubscriptions(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all studentsSubscriptions")
	}
}
