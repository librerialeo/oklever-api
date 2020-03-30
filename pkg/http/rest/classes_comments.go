package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitClassesCommentsHandler initialize ClassesComments router
func InitClassesCommentsHandler(r *router.Router) {
	r.GET("/:id/comments", getAllClassesComments)
}

func getAllClassesComments(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all ClassesComments")
}
