package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitPurchasesHandler initialize purchases router
func InitPurchasesHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllPurchases(s))
}

func getAllPurchases(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all purchases")
	}
}
