package handlers

import (
	"log"
	"net/http"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/chat/db"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/chat/models"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/chat/socket/client"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

type Repository struct {
	Hub      *models.Hub
	DB       db.Repository
	upgrader websocket.Upgrader
}

func NewHandler(hub *models.Hub, db db.Repository) *Repository {
	return &Repository{
		Hub: hub,
		DB:  db,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
	}
}

func (repo *Repository) Chat(w http.ResponseWriter, r *http.Request) {
	conn, err := repo.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	id1 := chi.URLParam(r, "id1")
	id2 := chi.URLParam(r, "id2")
	chatID := id1 + "-" + id2
	user := &models.Client{Hub: repo.Hub, Conn: conn, Send: make(chan models.Message), ChatID: chatID}
	user.Hub.Connect <- user

	go client.Write(user)
	go client.Read(user, repo.DB)
}
