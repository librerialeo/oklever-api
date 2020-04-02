package rest

import (
	"github.com/savsgio/atreugo"
)

func sessionMiddleware(ctx *atreugo.RequestCtx) error {
	ctx.SetUserValue("asdf", "variable de sessión")
	return ctx.Next()
}
