package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitTeachersTeachingInstitutionsHandler initialize TeachersTeachingInstitutions router
func InitTeachersTeachingInstitutionsHandler(r *router.Router) {
	r.GET("/", getAllTeachersTeachingInstitutions)
}

func getAllTeachersTeachingInstitutions(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all TeachersTeachingInstitutions")
}
