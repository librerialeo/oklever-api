package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitCoursesHandler initialize Courses router
func InitCoursesHandler(r *router.Router) {
	r.GET("/", getAllCourses)
}

func getAllCourses(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all Courses")
}
