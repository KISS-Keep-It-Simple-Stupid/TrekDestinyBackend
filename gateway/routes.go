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
	mux.Get("/get-offer", handler.GetOffer)
	mux.Get("/get-card-profile", handler.GetCardProfile)
	mux.Post("/upload", handler.UploadImage)
	mux.Get("/public-profile/{username}", handler.PublicProfile)
	mux.Post("/create-post", handler.CreatePost)
	mux.Get("/get-my-post", handler.GetMyPost)
	mux.Post("/get-post-host", handler.GetPostHost)
	mux.Post("/accept-offer", handler.AcceptOffer)
	mux.Post("/reject-offer", handler.RejectOffer)
	mux.Post("/public-profile-host", handler.PublicProfileHost)
	mux.Post("/chat-list" , handler.ChatListPost)
	mux.Get("/chat-list" , handler.ChatList)
	mux.Post("/edit-announcement", handler.EditAnnouncement)
	mux.Post("/delete-announcement", handler.DeleteAnnouncement)
	mux.Post("/edit-post", handler.EditPost)
	mux.Post("/host-info", handler.HostInfo)
	return mux
}
