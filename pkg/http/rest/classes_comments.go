package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo/v11"
)

// InitClassesCommentsHandler initialize ClassesComments router
func InitClassesCommentsHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllClassesComments(s))
}

func getAllClassesComments(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all ClassesComments")
		return err
	}
}
