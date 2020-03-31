package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitStudentsSubscriptionsHandler initialize studentsSubscriptions router
func InitStudentsSubscriptionsHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllStudentsSubscriptions(s))
}

func getAllStudentsSubscriptions(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all studentsSubscriptions")
		return err
	}
}
