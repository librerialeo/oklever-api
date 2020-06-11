package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo/v11"
)

// InitPurchasesHandler initialize purchases router
func InitPurchasesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllPurchases(s))
}

func getAllPurchases(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all purchases")
		return err
	}
}
