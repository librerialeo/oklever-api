package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitCoursesDatasheetsHandler initialize coursesDatasheets router
func InitCoursesDatasheetsHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllCoursesDatasheets(s))
}

func getAllCoursesDatasheets(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all coursesDatasheets")
	}
}
