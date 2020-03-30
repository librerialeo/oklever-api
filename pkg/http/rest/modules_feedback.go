package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitModulesFeedbackHandler initialize ModulesFeedback router
func InitModulesFeedbackHandler(r *router.Router) {
	r.GET("/", getAllModulesFeedback)
}

func getAllModulesFeedback(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all ModulesFeedback")
}
