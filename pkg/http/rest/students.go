package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitStudentsHandler initialize Students router
func InitStudentsHandler(r *router.Router) {
	r.GET("/", getAllStudents)
}

func getAllStudents(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all Students")
}
