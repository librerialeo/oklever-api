package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitCoursesLevelsHandler initialize coursesLevels router
func InitCoursesLevelsHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllCoursesLevels(s))
}

func getAllCoursesLevels(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all coursesLevels")
		return err
	}
}
