package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitUsersInsigniasHandler initialize usersInsignias router
func InitUsersInsigniasHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllUsersInsignias(s))
}

func getAllUsersInsignias(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all usersInsignias")
		return err
	}
}
