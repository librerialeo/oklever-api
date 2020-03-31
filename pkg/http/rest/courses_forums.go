package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitCoursesForumsHandler initialize coursesForums router
func InitCoursesForumsHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllCoursesForums(s))
}

func getAllCoursesForums(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all coursesForums")
		return err
	}
}
