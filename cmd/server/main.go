package main
import (
	"context"
	"log"

	"github.com/librerialeo/oklever-api/pkg/http/rest"
	"github.com/valyala/fasthttp"

	"github.com/fasthttp/router"
	"github.com/jackc/pgx"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "database=oklever")
	if err != nil {
		log.Fatal(err)
	}
	router := router.New()
	rest.InitRouterHandler(router, conn)
	log.Fatal(fasthttp.ListenAndServe(":8010", router.Handler))
}
