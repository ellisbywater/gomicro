package main

import (
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
}

func main() {
	app := Config{}

	log.Println("Starting broker-service on port %s", webPort)

	// define http server
	srv := &http.Server{
		Addr:    ":" + webPort,
		Handler: app.routes(),
	}

	// start http server
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
