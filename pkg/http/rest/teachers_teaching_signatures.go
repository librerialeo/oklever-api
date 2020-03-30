package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitTeachersTeachingSignaturesHandler initialize TeachersTeachingSignatures router
func InitTeachersTeachingSignaturesHandler(r *router.Router) {
	r.GET("/", getAllTeachersTeachingSignatures)
}

func getAllTeachersTeachingSignatures(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all TeachersTeachingSignatures")
}
