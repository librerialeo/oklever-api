package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitTeachersResearchHandler initialize TeachersResearch router
func InitTeachersResearchHandler(r *router.Router) {
	r.GET("/", getAllTeachersResearch)
}

func getAllTeachersResearch(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all TeachersResearch")
}
