package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitLanguagesHandler initialize languages router
func InitLanguagesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllLanguages(s))
}

func getAllLanguages(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all languages")
		return err
	}
}
