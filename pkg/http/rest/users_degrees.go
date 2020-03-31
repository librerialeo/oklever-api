package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitUsersDegreesHandler initialize usersDegrees router
func InitUsersDegreesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllUsersDegrees(s))
}

func getAllUsersDegrees(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all usersDegrees")
		return err
	}
}
