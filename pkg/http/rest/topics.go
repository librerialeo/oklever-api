package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitTopicsHandler initialize topics router
func InitTopicsHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllTopics(s))
}

func getAllTopics(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all topics")
	}
}
