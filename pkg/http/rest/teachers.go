package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitTeachersHandler initialize teachers router
func InitTeachersHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllTeachers(s))
}

func getAllTeachers(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all teachers")
	}
}
