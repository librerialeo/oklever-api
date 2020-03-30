package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitPurchasesProductsHandler initialize purchasesProducts router
func InitPurchasesProductsHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllPurchasesProducts(s))
}

func getAllPurchasesProducts(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all purchasesProducts")
	}
}
