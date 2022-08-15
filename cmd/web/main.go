package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jacstn/golang-url-shortner/config"
	"github.com/jacstn/golang-url-shortner/pkg/handlers"
)

const portNumber = ":3000"

var app = config.AppConfig{
	Production: false,
}

func main() {

	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = app.Production

	app.Session = session
	handlers.NewHandlers(&app)
	fmt.Println(fmt.Sprintf("Starting application %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal("Cannot start server")
	}
}
