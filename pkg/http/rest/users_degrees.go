package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitUsersDegreesHandler initialize usersDegrees router
func InitUsersDegreesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllUsersDegrees(s))
}

func getAllUsersDegrees(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all usersDegrees")
	}
}
