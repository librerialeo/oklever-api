package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitClassesCommentsHandler initialize ClassesComments router
func InitClassesCommentsHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllClassesComments(s))
}

func getAllClassesComments(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all ClassesComments")
	}
}
