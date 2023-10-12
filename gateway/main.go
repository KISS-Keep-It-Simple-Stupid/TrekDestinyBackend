package main

import (
	"log"
	"net/http"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/gateway/handlers"
)

const port string = ":9090"

func main() {
	handler := handlers.New()
	routes := getRoutes(handler)
	server := http.Server{
		Addr:    port,
		Handler: routes,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
