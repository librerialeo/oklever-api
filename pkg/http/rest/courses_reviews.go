package rest

import (
	"github.com/fasthttp/router"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/valyala/fasthttp"
)

// InitCoursesReviewsHandler initialize coursesReviews router
func InitCoursesReviewsHandler(r *router.Router, s *service.Service) {
	r.GET("/", getAllCoursesReviews(s))
}

func getAllCoursesReviews(s *service.Service) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("get all coursesReviews")
	}
}
