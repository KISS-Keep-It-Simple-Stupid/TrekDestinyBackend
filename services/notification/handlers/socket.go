package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/client"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/models"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

type Repository struct {
	upgrader websocket.Upgrader
	hub      *models.Hub
}

func New(h *models.Hub) *Repository {
	return &Repository{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
		hub: h,
	}
}

func (repo *Repository) Subscribe(w http.ResponseWriter, r *http.Request) {
	conn, err := repo.upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("content/type", "application/json")
		errorMessage := struct {
			Message string `json:"message"`
		}{
			Message: "can not open websocket connection",
		}
		errResponse, _ := json.Marshal(errorMessage)
		w.Write(errResponse)
		return
	}
	user_id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	c := &models.Client{
		ID:          user_id,
		SendChannel: make(chan models.Signal),
		Conn:        conn,
	}
	repo.hub.Connect <- c
	go client.Read(c, repo.hub)
	go client.Write(c)
}
