package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitSubscriptionsHandler initialize subscriptions router
func InitSubscriptionsHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllSubscriptions(s))
}

func getAllSubscriptions(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all subscriptions")
		return err
	}
}
