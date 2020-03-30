package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitQuestionsHandler initialize Questions router
func InitQuestionsHandler(r *router.Router) {
	r.GET("/", getAllQuestions)
}

func getAllQuestions(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all Questions")
}
