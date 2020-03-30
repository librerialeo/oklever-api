package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitTopicsResourcesHandler initialize TopicsResources router
func InitTopicsResourcesHandler(r *router.Router) {
	r.GET("/", getAllTopicsResources)
}

func getAllTopicsResources(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all TopicsResources")
}
