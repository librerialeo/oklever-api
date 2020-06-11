package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo/v11"
)

// InitStudentsProjectsHistoryHandler initialize studentsProjectsHistory router
func InitStudentsProjectsHistoryHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllStudentsProjectsHistory(s))
}

func getAllStudentsProjectsHistory(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all studentsProjectsHistory")
		return err
	}
}
