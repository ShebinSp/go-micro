package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80" //8081

// This will be the receiver for the application
type Config struct{}

func main() {
	// variable of type config
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s",webPort),
		Handler: app.routes(),
	}

	// start the server
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}

}
