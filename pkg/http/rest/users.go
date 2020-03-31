package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitUsersHandler initialize users router
func InitUsersHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllUsers(s))
}

func getAllUsers(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all users")
		return err
	}
}
