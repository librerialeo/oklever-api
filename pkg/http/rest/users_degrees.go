package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitUsersDegreesHandler initialize UsersDegrees router
func InitUsersDegreesHandler(r *router.Router) {
	r.GET("/", getAllUsersDegrees)
}

func getAllUsersDegrees(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all UsersDegrees")
}
