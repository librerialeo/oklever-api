package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitQuestionsTypesHandler initialize QuestionsTypes router
func InitQuestionsTypesHandler(r *router.Router) {
	r.GET("/", getAllQuestionsTypes)
}

func getAllQuestionsTypes(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all QuestionsTypes")
}
