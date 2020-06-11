package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo/v11"
)

// InitUsersLanguagesHandler initialize usersLanguages router
func InitUsersLanguagesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllUsersLanguages(s))
}

func getAllUsersLanguages(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all users Languages")
		return err
	}
}
