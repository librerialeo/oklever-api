package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitTeachersResearchHandler initialize teachersResearch router
func InitTeachersResearchHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllTeachersResearch(s))
}

func getAllTeachersResearch(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all teachersResearch")
		return err
	}
}
