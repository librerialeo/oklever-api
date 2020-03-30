package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitStudentsQuizzesHandler initialize StudentsQuizzes router
func InitStudentsQuizzesHandler(r *router.Router) {
	r.GET("/", getAllStudentsQuizzes)
}

func getAllStudentsQuizzes(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all StudentsQuizzes")
}
