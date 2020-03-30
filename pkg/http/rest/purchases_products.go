package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitPurchasesProductsHandler initialize PurchasesProducts router
func InitPurchasesProductsHandler(r *router.Router) {
	r.GET("/", getAllPurchasesProducts)
}

func getAllPurchasesProducts(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all PurchasesProducts")
}
