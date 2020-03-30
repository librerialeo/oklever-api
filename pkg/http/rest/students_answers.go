package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitStudentsAnswersHandler initialize studentsAnswers router
func InitStudentsAnswersHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllStudentsAnswers(s))
}

func getAllStudentsAnswers(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all studentsAnswers")
	}
}
