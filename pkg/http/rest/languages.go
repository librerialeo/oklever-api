package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitLanguagesHandler initialize Languages router
func InitLanguagesHandler(r *router.Router) {
	r.GET("/", getAllLanguages)
}

func getAllLanguages(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all Languages")
}
