package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitTeachersTeachingInstitutionsHandler initialize teachersTeachingInstitutions router
func InitTeachersTeachingInstitutionsHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllTeachersTeachingInstitutions(s))
}

func getAllTeachersTeachingInstitutions(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all teachersTeachingInstitutions")
		return err
	}
}
