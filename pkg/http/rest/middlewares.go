package rest

import (
	"github.com/savsgio/atreugo/v11"
)

func sessionMiddleware(ctx *atreugo.RequestCtx) error {
	ctx.SetUserValue("asdf", "variable de sessión")
	return ctx.Next()
}
