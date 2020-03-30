package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitUsersInsigniasHandler initialize UsersInsignias router
func InitUsersInsigniasHandler(r *router.Router) {
	r.GET("/", getAllUsersInsignias)
}

func getAllUsersInsignias(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all UsersInsignias")
}
