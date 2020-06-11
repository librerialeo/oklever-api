package rest

import (
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo/v11"
)

// InitTopicsHandler initialize topics router
func InitTopicsHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllTopics(s))
}

func getAllTopics(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all topics")
		return err
	}
}
