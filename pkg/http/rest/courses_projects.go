package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitCoursesProjectsHandler initialize CoursesProjects router
func InitCoursesProjectsHandler(r *router.Router) {
	r.GET("/", getAllCoursesProjects)
}

func getAllCoursesProjects(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all CoursesProjects")
}
