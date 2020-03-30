package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitStudentsHandler initialize students router
func InitStudentsHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllStudents(s))
}

func getAllStudents(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all students")
	}
}
