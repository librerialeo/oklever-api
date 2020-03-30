package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitCoursesLevelsHandler initialize coursesLevels router
func InitCoursesLevelsHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllCoursesLevels(s))
}

func getAllCoursesLevels(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all coursesLevels")
	}
}
