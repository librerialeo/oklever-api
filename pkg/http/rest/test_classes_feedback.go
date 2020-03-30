package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitTestClassesFeedbackHandler initialize testClassesFeedback router
func InitTestClassesFeedbackHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllTestClassesFeedback(s))
}

func getAllTestClassesFeedback(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all testClassesFeedback")
	}
}
