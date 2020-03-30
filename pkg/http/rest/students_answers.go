package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitStudentsAnswersHandler initialize StudentsAnswers router
func InitStudentsAnswersHandler(r *router.Router) {
	r.GET("/", getAllStudentsAnswers)
}

func getAllStudentsAnswers(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all StudentsAnswers")
}
