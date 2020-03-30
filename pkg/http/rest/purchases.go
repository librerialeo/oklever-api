package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitPurchasesHandler initialize Purchases router
func InitPurchasesHandler(r *router.Router) {
	r.GET("/", getAllPurchases)
}

func getAllPurchases(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all Purchases")
}
