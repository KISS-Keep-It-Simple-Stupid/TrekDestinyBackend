package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/chat/db"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/chat/helper"
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
	id3 := chi.URLParam(r, "id3")
	chatID := id1 + "-" + id2 + "-" +id3
	user := &models.Client{Hub: repo.Hub, Conn: conn, Send: make(chan models.Message), ChatID: chatID}
	user.Hub.Connect <- user

	go client.Write(user)
	go client.Read(user, repo.DB)
}

func (repo *Repository) Count(w http.ResponseWriter, r *http.Request) {
	id1 := chi.URLParam(r, "id1")
	id2 := chi.URLParam(r, "id2")
	id3 := chi.URLParam(r, "id3")
	chatID := id1 + "-" + id2 + "-" +id3
	count := repo.DB.GetCount(chatID)
	helper.ResponseGenerator(w, struct {
		Count int `json:"count"`
	}{
		Count: count,
	})
}

func (repo *Repository) History(w http.ResponseWriter, r *http.Request) {
	id1 := chi.URLParam(r, "id1")
	id2 := chi.URLParam(r, "id2")
	id3 := chi.URLParam(r, "id3")
	chatID := id1 + "-" + id2 + "-" +id3
	pg := r.URL.Query().Get("page")
	page := 1
	var err error
	if pg != "" {
		page, err = strconv.Atoi(pg)
		if err != nil {
			helper.MessageGenerator(w, "page type must be integer", http.StatusBadRequest)
			return
		}
	}
	countInQuery := r.URL.Query().Get("count")
	count := 0
	if countInQuery != "" {
		count, err = strconv.Atoi(countInQuery)
		if err != nil {
			helper.MessageGenerator(w, "count type must be integer", http.StatusBadRequest)
			return
		}
	}
	response, err := repo.DB.GetHistory(chatID, page, count)
	if err != nil {
		log.Println(err)
		helper.MessageGenerator(w, "internal error while getting chat history - chat service", http.StatusInternalServerError)
		return
	}

	helper.ResponseGenerator(w, response)
}
