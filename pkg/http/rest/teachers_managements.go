package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitTeachersManagementsHandler initialize teachersManagements router
func InitTeachersManagementsHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllTeachersManagements(s))
}

func getAllTeachersManagements(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all teachersManagements")
	}
}
