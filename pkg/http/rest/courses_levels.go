package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitCoursesLevelsHandler initialize CoursesLevels router
func InitCoursesLevelsHandler(r *router.Router) {
	r.GET("/", getAllCoursesLevels)
}

func getAllCoursesLevels(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all CoursesLevels")
}
