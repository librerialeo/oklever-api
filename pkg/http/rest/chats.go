package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitChatsHandler initialize Chats router
func InitChatsHandler(r *router.Router) {
	r.GET("/", getAllChats)
}

func getAllChats(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all chats")
}
