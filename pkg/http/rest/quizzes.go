package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitQuizzesHandler initialize quizzes router
func InitQuizzesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllQuizzes(s))
}

func getAllQuizzes(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all quizzes")
	}
}
