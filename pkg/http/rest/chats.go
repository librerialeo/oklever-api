package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitChatsHandler initialize chats router
func InitChatsHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllChats(s))
}

func getAllChats(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all chats")
	}
}
