package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitStudentsProjectsHistoryHandler initialize StudentsProjectsHistory router
func InitStudentsProjectsHistoryHandler(r *router.Router) {
	r.GET("/", getAllStudentsProjectsHistory)
}

func getAllStudentsProjectsHistory(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all StudentsProjectsHistory")
}
