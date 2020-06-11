package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo/v11"
)

// InitCoursesForumsCommentsHandler initialize coursesForumsComments router
func InitCoursesForumsCommentsHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllCoursesForumsComments(s))
}

func getAllCoursesForumsComments(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all coursesForumsComments")
		return err
	}
}
