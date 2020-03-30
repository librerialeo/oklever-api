package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitCoursesHandler initialize courses router
func InitCoursesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllCourses(s))
}

func getAllCourses(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all courses")
	}
}
