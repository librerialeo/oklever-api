package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitQuestionsHandler initialize questions router
func InitQuestionsHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllQuestions(s))
}

func getAllQuestions(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all questions")
	}
}
