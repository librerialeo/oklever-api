package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitQuestionsOptionsHandler initialize questionsOptions router
func InitQuestionsOptionsHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllQuestionsOptions(s))
}

func getAllQuestionsOptions(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all questionsOptions")
	}
}
