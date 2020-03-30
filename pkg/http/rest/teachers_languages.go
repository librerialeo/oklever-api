package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitTeachersLanguagesHandler initialize TeachersLanguages router
func InitTeachersLanguagesHandler(r *router.Router) {
	r.GET("/", getAllTeachersLanguages)
}

func getAllTeachersLanguages(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all TeachersLanguages")
}
