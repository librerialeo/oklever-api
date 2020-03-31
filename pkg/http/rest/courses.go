package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitCoursesHandler initialize courses router
func InitCoursesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllCourses(s))
}

func getAllCourses(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all courses")
		return err
	}
}
