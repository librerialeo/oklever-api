package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitTeachersTeachingSignaturesHandler initialize teachersTeachingSignatures router
func InitTeachersTeachingSignaturesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllTeachersTeachingSignatures(s))
}

func getAllTeachersTeachingSignatures(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all teachersTeachingSignatures")
	}
}
