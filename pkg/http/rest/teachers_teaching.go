package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitTeachersTeachingHandler initialize TeachersTeaching router
func InitTeachersTeachingHandler(r *router.Router) {
	r.GET("/", getAllTeachersTeaching)
}

func getAllTeachersTeaching(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all TeachersTeaching")
}
