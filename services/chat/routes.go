package main

import (
	"net/http"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/chat/handlers"
	"github.com/go-chi/chi/v5"
)

func GetRoutes(handler *handlers.Repository) http.Handler {
	mux := chi.NewRouter()
	mux.Get("/{id1}/{id2}", handler.Chat)
	mux.Get("/count/{id1}/{id2}", handler.Count)
	mux.Get("/history/{id1}/{id2}", handler.History)
	return mux
}
