package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitCoursesReviewsHandler initialize CoursesReviews router
func InitCoursesReviewsHandler(r *router.Router) {
	r.GET("/", getAllCoursesReviews)
}

func getAllCoursesReviews(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all CoursesReviews")
}
