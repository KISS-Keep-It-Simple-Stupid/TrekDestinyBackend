package main

import (
	"net/http"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/gateway/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func getRoutes(handler *handlers.Repository) http.Handler {
	mux := chi.NewRouter()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	mux.Post("/signup", handler.SignUp)
	mux.Post("/login", handler.Login)
	mux.Post("/refresh", handler.Refresh)
	mux.Post("/forget-password", handler.ForgetPassword)
	mux.Post("/reset-password", handler.ResetPassword)
	mux.Get("/verify-email", handler.EmailVerification)
	mux.Get("/profile", handler.Profile)
	mux.Post("/edit-profile", handler.EditProfile)
	mux.Post("/create-card", handler.CreateCard)
	mux.Get("/get-card", handler.GetCard)
	mux.Post("/create-offer", handler.CreateOffer)
	mux.Post("/get-offer", handler.GetOffer)
	mux.Get("/get-card-profile", handler.GetCardProfile)
	mux.Post("/upload", handler.UploadImage)
	return mux
}
