package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo/v11"
)

// InitStudentsCoursesHandler initialize studentsCourses router
func InitStudentsCoursesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllStudentsCourses(s))
}

func getAllStudentsCourses(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all studentsCourses")
		return err
	}
}
