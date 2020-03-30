package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitCoursesForumsHandler initialize coursesForums router
func InitCoursesForumsHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllCoursesForums(s))
}

func getAllCoursesForums(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all coursesForums")
	}
}
