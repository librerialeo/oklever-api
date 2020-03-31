package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitCoursesDatasheetsHandler initialize coursesDatasheets router
func InitCoursesDatasheetsHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllCoursesDatasheets(s))
}

func getAllCoursesDatasheets(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all coursesDatasheets")
		return err
	}
}
