package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitTeachersTeachingHandler initialize teachersTeaching router
func InitTeachersTeachingHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllTeachersTeaching(s))
}

func getAllTeachersTeaching(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all teachersTeaching")
	}
}
