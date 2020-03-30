package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitTeachersManagementsHandler initialize TeachersManagements router
func InitTeachersManagementsHandler(r *router.Router) {
	r.GET("/", getAllTeachersManagements)
}

func getAllTeachersManagements(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all TeachersManagements")
}
