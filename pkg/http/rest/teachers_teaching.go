package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitTeachersTeachingHandler initialize teachersTeaching router
func InitTeachersTeachingHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllTeachersTeaching(s))
}

func getAllTeachersTeaching(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all teachersTeaching")
		return err
	}
}
