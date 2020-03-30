package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitUsersHandler initialize users router
func InitUsersHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllUsers(s))
}

func getAllUsers(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all users")
	}
}
