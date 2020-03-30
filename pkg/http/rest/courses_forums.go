package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitCoursesForumsHandler initialize CoursesForums router
func InitCoursesForumsHandler(r *router.Router) {
	r.GET("/", getAllCoursesForums)
}

func getAllCoursesForums(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all CoursesForums")
}
