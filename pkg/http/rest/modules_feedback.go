package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitModulesFeedbackHandler initialize modulesFeedback router
func InitModulesFeedbackHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllModulesFeedback(s))
}

func getAllModulesFeedback(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all modulesFeedback")
	}
}
