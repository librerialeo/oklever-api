package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitStudentsProjectsHandler initialize StudentsProjects router
func InitStudentsProjectsHandler(r *router.Router) {
	r.GET("/", getAllStudentsProjects)
}

func getAllStudentsProjects(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all StudentsProjects")
}
