package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitStudentsProjectsHandler initialize studentsProjects router
func InitStudentsProjectsHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllStudentsProjects(s))
}

func getAllStudentsProjects(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all studentsProjects")
	}
}
