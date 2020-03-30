package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitInsigniasHandler initialize Insignias router
func InitInsigniasHandler(r *router.Router) {
	r.GET("/", getAllInsignias)
}

func getAllInsignias(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all Insignias")
}
