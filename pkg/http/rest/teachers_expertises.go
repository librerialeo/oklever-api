package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitTeachersExpertisesHandler initialize teachersExpertises router
func InitTeachersExpertisesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllTeachersExpertises(s))
}

func getAllTeachersExpertises(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all teachersExpertises")
		return err
	}
}
