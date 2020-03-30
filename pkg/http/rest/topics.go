package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitTopicsHandler initialize Topics router
func InitTopicsHandler(r *router.Router) {
	r.GET("/", getAllTopics)
}

func getAllTopics(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all Topics")
}
