package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitTeachersHandler initialize teachers router
func InitTeachersHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllTeachers(s))
}

func getAllTeachers(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all teachers")
		return err
	}
}
