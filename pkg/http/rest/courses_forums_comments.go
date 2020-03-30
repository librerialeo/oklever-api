package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitCoursesForumsCommentsHandler initialize CoursesForumsComments router
func InitCoursesForumsCommentsHandler(r *router.Router) {
	r.GET("/", getAllCoursesForumsComments)
}

func getAllCoursesForumsComments(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all CoursesForumsComments")
}
