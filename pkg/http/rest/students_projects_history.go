package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitStudentsProjectsHistoryHandler initialize studentsProjectsHistory router
func InitStudentsProjectsHistoryHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllStudentsProjectsHistory(s))
}

func getAllStudentsProjectsHistory(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all studentsProjectsHistory")
	}
}
