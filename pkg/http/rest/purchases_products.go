package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo/v11"
)

// InitPurchasesProductsHandler initialize purchasesProducts router
func InitPurchasesProductsHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllPurchasesProducts(s))
}

func getAllPurchasesProducts(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all purchasesProducts")
		return err
	}
}
