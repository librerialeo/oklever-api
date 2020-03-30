package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitTeachersTeachingInstitutionsHandler initialize teachersTeachingInstitutions router
func InitTeachersTeachingInstitutionsHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllTeachersTeachingInstitutions(s))
}

func getAllTeachersTeachingInstitutions(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all teachersTeachingInstitutions")
	}
}
