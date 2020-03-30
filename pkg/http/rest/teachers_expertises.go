package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitTeachersExpertisesHandler initialize TeachersExpertises router
func InitTeachersExpertisesHandler(r *router.Router) {
	r.GET("/", getAllTeachersExpertises)
}

func getAllTeachersExpertises(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all TeachersExpertises")
}
