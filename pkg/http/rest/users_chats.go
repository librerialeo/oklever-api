package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitUsersChatsHandler initialize UsersChats router
func InitUsersChatsHandler(r *router.Router) {
	r.GET("/", getAllUsersChats)
}

func getAllUsersChats(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all UsersChats")
}
