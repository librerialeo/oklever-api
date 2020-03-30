package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitUsersInsigniasHandler initialize usersInsignias router
func InitUsersInsigniasHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllUsersInsignias(s))
}

func getAllUsersInsignias(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all usersInsignias")
	}
}
