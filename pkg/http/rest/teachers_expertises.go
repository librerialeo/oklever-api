package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitTeachersExpertisesHandler initialize teachersExpertises router
func InitTeachersExpertisesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllTeachersExpertises(s))
}

func getAllTeachersExpertises(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all teachersExpertises")
	}
}
