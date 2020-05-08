package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitUsersAcademyHandler initialize usersAcademy router
func InitUsersAcademyHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllUsersAcademy(s))
}

func getAllUsersAcademy(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all users academy")
		return err
	}
}
