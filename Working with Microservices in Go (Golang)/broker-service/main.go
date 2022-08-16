package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8080"

func main() {
	app := Config{
		domain: "localhost",
	}

	log.Printf("Starting broker service on port %s\n", webPort)
	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}