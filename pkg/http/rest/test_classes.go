package rest

import (
	"fmt"

	"github.com/librerialeo/oklever-api/pkg/service"
	"github.com/savsgio/atreugo"
)

// InitTestClassesHandler initialize testClasses router
func InitTestClassesHandler(r *atreugo.Router, s *service.Service) {
	r.GET("/", getAllTestClasses(s))
	r.POST("/add", addTeachersTestClass(s))
}

func getAllTestClasses(s *service.Service) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		_, err := ctx.WriteString("get all testClasses")
		return err
	}
}

// addTeahcersTestClass teachers add new test class
func addTeachersTestClass(s *service.Service) atreugo.View {
	fmt.Println(s)
	return nil
}
