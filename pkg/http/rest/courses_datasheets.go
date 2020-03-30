package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitCoursesDatasheetsHandler initialize CoursesDatasheets router
func InitCoursesDatasheetsHandler(r *router.Router) {
	r.GET("/", getAllCoursesDatasheets)
}

func getAllCoursesDatasheets(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all CoursesDatasheets")
}
