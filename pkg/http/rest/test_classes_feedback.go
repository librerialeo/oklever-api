package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitTestClassesFeedbackHandler initialize TestClassesFeedback router
func InitTestClassesFeedbackHandler(r *router.Router) {
	r.GET("/", getAllTestClassesFeedback)
}

func getAllTestClassesFeedback(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all TestClassesFeedback")
}
