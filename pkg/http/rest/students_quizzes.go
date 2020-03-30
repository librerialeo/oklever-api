package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitStudentsQuizzesHandler initialize studentsQuizzes router
func InitStudentsQuizzesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllStudentsQuizzes(s))
}

func getAllStudentsQuizzes(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all studentsQuizzes")
	}
}
