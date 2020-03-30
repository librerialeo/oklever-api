package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitTeachersResearchHandler initialize teachersResearch router
func InitTeachersResearchHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllTeachersResearch(s))
}

func getAllTeachersResearch(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all teachersResearch")
	}
}
