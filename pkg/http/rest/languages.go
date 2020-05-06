package rest

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitLanguagesHandler initialize languages router
func InitLanguagesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllLanguages(s))
}

func getAllLanguages(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		languages, err := s.GetAllLanguages()
		response := ServerResponse{
			Error: nil,
			Token: ctx.UserValue("token"),
			Data:  "",
		}
		if err != nil {
			response.Error = err
			message, err := jsoniter.Marshal(response)
			if err != nil {
				return err
			}
			_, err = ctx.Write(message)
			return err
		}
		response.Data = languages
		message, err := jsoniter.Marshal(response)
		if err != nil {
			return err
		}
		_, err = ctx.Write(message)
		return err
	}
}
