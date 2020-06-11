package main

import (
	"fmt"
	"log"
	"os"

	"github.com/savsgio/atreugo/v11"

	"github.com/librerialeo/oklever-api/pkg/http/rest"

	"github.com/jackc/pgx"
)

func main() {
	conn, err := pgx.Connect(pgx.ConnConfig{
		Database: os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatal(err)
	}
	config := atreugo.Config{
		Addr: fmt.Sprintf("0.0.0.0:%s", os.Getenv("OKLEVER_PORT")),
	}
	router := atreugo.New(config)
	rest.InitRouterHandler(router, conn)
	log.Fatal(router.ListenAndServe())
}
