package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitRolesHandler initialize Roles router
func InitRolesHandler(r *router.Router) {
	r.GET("/", getAllRoles)
}

func getAllRoles(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all Roles")
}
