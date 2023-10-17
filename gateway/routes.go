package main

import (
	"net/http"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/gateway/handlers"
	"github.com/go-chi/chi/v5"
)

func getRoutes(handler *handlers.Repository) http.Handler {
	mux := chi.NewRouter()
	mux.Post("/signup", handler.SignUp)
	mux.Post("/login", handler.Login)
	mux.Post("/refresh", handler.Refresh)
	mux.Post("/forget-password",handler.ForgetPassword)
	return mux
}
