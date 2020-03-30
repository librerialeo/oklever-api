package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitTeachersLanguagesHandler initialize teachersLanguages router
func InitTeachersLanguagesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllTeachersLanguages(s))
}

func getAllTeachersLanguages(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all teachersLanguages")
	}
}
