package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo/v11"
)

// InitLanguagesHandler initialize languages router
func InitLanguagesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllLanguages(s))
}

func getAllLanguages(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		languages, err := s.GetAllLanguages()
		if err != nil {
			ctx.SetUserValue("error", err)
			return SendResponse(ctx)
		}
		return SendResponse(ctx, languages)
	}
}
