package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitCoursesProjectsHandler initialize coursesProjects router
func InitCoursesProjectsHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllCoursesProjects(s))
}

func getAllCoursesProjects(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all coursesProjects")
		return err
	}
}
