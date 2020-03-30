package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitUsersHandler initialize Users router
func InitUsersHandler(r *router.Router) {
	r.GET("/", getAllUsers)
}

func getAllUsers(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all Users")
}
