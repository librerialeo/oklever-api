package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitChatsHandler initialize chats router
func InitChatsHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllChats(s)).UseBefore(sessionMiddleware)
}

func getAllChats(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		logger := ctx.Logger()
		asdf := ctx.UserValue("asdf").(string)
		logger.Printf("%s", asdf)
		_, err := ctx.WriteString("get all chats")
		return err
	}
}
