package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitCoursesForumsCommentsHandler initialize coursesForumsComments router
func InitCoursesForumsCommentsHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllCoursesForumsComments(s))
}

func getAllCoursesForumsComments(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all coursesForumsComments")
	}
}
