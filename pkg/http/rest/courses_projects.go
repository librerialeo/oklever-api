package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitCoursesProjectsHandler initialize coursesProjects router
func InitCoursesProjectsHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllCoursesProjects(s))
}

func getAllCoursesProjects(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all coursesProjects")
	}
}
