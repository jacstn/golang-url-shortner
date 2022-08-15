package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":3000"

func main() {
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
