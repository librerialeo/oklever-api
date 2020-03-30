package rest

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// InitQuestionsOptionsHandler initialize QuestionsOptions router
func InitQuestionsOptionsHandler(r *router.Router) {
	r.GET("/", getAllQuestionsOptions)
}

func getAllQuestionsOptions(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("get all QuestionsOptions")
}
