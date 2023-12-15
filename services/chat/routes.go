package main

import (
	"net/http"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/chat/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func GetRoutes(handler *handlers.Repository) http.Handler {
	mux := chi.NewRouter()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	mux.Get("/{id1}/{id2}/{id3}", handler.Chat)
	mux.Get("/count/{id1}/{id2}/{id3}", handler.Count)
	mux.Get("/history/{id1}/{id2}/{id3}", handler.History)
	return mux
}
