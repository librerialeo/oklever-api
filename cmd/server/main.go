package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/savsgio/atreugo"

	"github.com/librerialeo/oklever-api/pkg/http/rest"

	"github.com/jackc/pgx"
)

func main() {
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("database=%s", os.Getenv("DB_NAME")))
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

// package main

// import (
// 	"encoding/base64"
// 	"fmt"
// 	"log"
// 	"strings"

// 	"github.com/fasthttp/router"
// 	"github.com/elithrar/simple-scrypt"
// 	"github.com/valyala/fasthttp"
// )

// // basicAuth returns the username and password provided in the request's
// // Authorization header, if the request uses HTTP Basic Authentication.
// // See RFC 2617, Section 2.
// func basicAuth(ctx *fasthttp.RequestCtx) (username, password string, ok bool) {
// 	auth := ctx.Request.Header.Peek("Authorization")
// 	if auth == nil {
// 		return
// 	}
// 	return parseBasicAuth(string(auth))
// }

// // parseBasicAuth parses an HTTP Basic Authentication string.
// // "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==" returns ("Aladdin", "open sesame", true).
// func parseBasicAuth(auth string) (username, password string, ok bool) {
// 	const prefix = "Basic "
// 	if !strings.HasPrefix(auth, prefix) {
// 		return
// 	}
// 	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
// 	if err != nil {
// 		return
// 	}
// 	cs := string(c)
// 	s := strings.IndexByte(cs, ':')
// 	if s < 0 {
// 		return
// 	}
// 	return cs[:s], cs[s+1:], true
// }

// // BasicAuth is the basic auth handler
// func BasicAuth(h fasthttp.RequestHandler, requiredUser string, requiredPasswordHash []byte) fasthttp.RequestHandler {
// 	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
// 		// Get the Basic Authentication credentials
// 		user, password, hasAuth := basicAuth(ctx)

// 		// WARNING:
// 		// DO NOT use plain-text passwords for real apps.
// 		// A simple string comparison using == is vulnerable to a timing attack.
// 		// Instead, use the hash comparison function found in your hash library.
// 		// This example uses scrypt, which is a solid choice for secure hashing:
// 		//   go get -u github.com/elithrar/simple-scrypt

// 		if hasAuth && user == requiredUser {

// 			// Uses the parameters from the existing derived key. Return an error if they don't match.
// 			err := scrypt.CompareHashAndPassword(requiredPasswordHash, []byte(password))

// 		    	if err != nil {

// 				// log error and request Basic Authentication again below.
// 				log.Fatal(err)

// 		    	} else {

// 				// Delegate request to the given handle
// 				h(ctx)
// 				return

// 		    	}

// 		}

// 		// Request Basic Authentication otherwise
// 		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusUnauthorized), fasthttp.StatusUnauthorized)
// 		ctx.Response.Header.Set("WWW-Authenticate", "Basic realm=Restricted")
// 	})
// }

// // Index is the index handler
// func Index(ctx *fasthttp.RequestCtx) {
// 	fmt.Fprint(ctx, "Not protected!\n")
// }

// // Protected is the Protected handler
// func Protected(ctx *fasthttp.RequestCtx) {
// 	fmt.Fprint(ctx, "Protected!\n")
// }

// func main() {
// 	user := "gordon"
// 	pass := "secret!"

// 	// generate a hashed password from the password above:
// 	hashedPassword, err := scrypt.GenerateFromPassword([]byte(pass), scrypt.DefaultParams)
//     	if err != nil {
//         	log.Fatal(err)
//     	}

// 	r := router.New()
// 	r.GET("/", Index)
// 	r.GET("/protected/", BasicAuth(Protected, user, hashedPassword))

// 	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
// }
