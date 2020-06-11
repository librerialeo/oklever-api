package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo/v11"
)

// InitUsersChatsHandler initialize usersChats router
func InitUsersChatsHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllUsersChats(s))
}

func getAllUsersChats(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all usersChats")
		return err
	}
}
