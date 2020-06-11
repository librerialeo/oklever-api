package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo/v11"
)

// InitCoursesReviewsHandler initialize coursesReviews router
func InitCoursesReviewsHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllCoursesReviews(s))
}

func getAllCoursesReviews(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all coursesReviews")
		return err
	}
}
