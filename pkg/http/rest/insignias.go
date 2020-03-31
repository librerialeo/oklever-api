package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitInsigniasHandler initialize insignias router
func InitInsigniasHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllInsignias(s))
}

func getAllInsignias(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all insignias")
		return err
	}
}
