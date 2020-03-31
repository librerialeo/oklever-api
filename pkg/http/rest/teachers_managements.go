package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitTeachersManagementsHandler initialize teachersManagements router
func InitTeachersManagementsHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllTeachersManagements(s))
}

func getAllTeachersManagements(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all teachersManagements")
		return err
	}
}
