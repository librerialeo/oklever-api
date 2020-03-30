package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitStudentsCoursesHandler initialize StudentsCourses router
func InitStudentsCoursesHandler(r *router.Router) {
	r.GET("/", getAllStudentsCourses)
}

func getAllStudentsCourses(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all StudentsCourses")
}
