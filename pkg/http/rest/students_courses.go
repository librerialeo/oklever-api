package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitStudentsCoursesHandler initialize studentsCourses router
func InitStudentsCoursesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllStudentsCourses(s))
}

func getAllStudentsCourses(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all studentsCourses")
	}
}
