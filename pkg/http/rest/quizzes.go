package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitQuizzesHandler initialize Quizzes router
func InitQuizzesHandler(r *router.Router) {
	r.GET("/", getAllQuizzes)
}

func getAllQuizzes(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all Quizzes")
}
