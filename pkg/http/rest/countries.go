package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitCountriesHandler initialize countries router
func InitCountriesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllCountries(s))
}

func getAllCountries(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all countries")
		return err
	}
}
