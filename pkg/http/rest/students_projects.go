package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitStudentsProjectsHandler initialize studentsProjects router
func InitStudentsProjectsHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllStudentsProjects(s))
}

func getAllStudentsProjects(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all studentsProjects")
		return err
	}
}
