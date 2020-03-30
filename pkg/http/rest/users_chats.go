package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitUsersChatsHandler initialize usersChats router
func InitUsersChatsHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllUsersChats(s))
}

func getAllUsersChats(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all usersChats")
	}
}
