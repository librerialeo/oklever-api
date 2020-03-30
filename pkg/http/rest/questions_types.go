package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitQuestionsTypesHandler initialize questionsTypes router
func InitQuestionsTypesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllQuestionsTypes(s))
}

func getAllQuestionsTypes(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all questionsTypes")
	}
}
