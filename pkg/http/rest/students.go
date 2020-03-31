package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitStudentsHandler initialize students router
func InitStudentsHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllStudents(s))
}

func getAllStudents(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all students")
		return err
	}
}
