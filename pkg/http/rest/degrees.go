package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitDegreesHandler initialize degrees router
func InitDegreesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllDegrees(s))
}

func getAllDegrees(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		degrees, err := s.GetAllDegrees()
		if err != nil {
			ctx.SetUserValue("error", err)
			return SendResponse(ctx)
		}
		return SendResponse(ctx, degrees)
	}
}
