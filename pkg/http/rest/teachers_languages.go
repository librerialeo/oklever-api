package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitTeachersLanguagesHandler initialize teachersLanguages router
func InitTeachersLanguagesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllTeachersLanguages(s))
}

func getAllTeachersLanguages(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all teachersLanguages")
		return err
	}
}
