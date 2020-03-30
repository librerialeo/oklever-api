package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitTeachersHandler initialize Teachers router
func InitTeachersHandler(r *router.Router) {
	r.GET("/", getAllTeachers)
}

func getAllTeachers(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all Teachers")
}
