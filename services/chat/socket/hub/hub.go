package hub

import (
	"log"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/chat/models"
)

func NewHub() *models.Hub {
	return &models.Hub{
		BroadCast:  make(chan models.Message),
		Connect:    make(chan *models.Client),
		Disconnect: make(chan *models.Client),
		Clients:    make(map[string]map[*models.Client]bool),
	}
}

func StartHub(h *models.Hub) {
	for {
		select {
		case client := <-h.Connect:
			if h.Clients[client.ChatID] == nil {
				h.Clients[client.ChatID] = make(map[*models.Client]bool)
			}
			h.Clients[client.ChatID][client] = true
			log.Println("a client connected to the chat:", client.ChatID)
		case client := <-h.Disconnect:
			if _, ok := h.Clients[client.ChatID][client]; ok {
				log.Println("a client disconnected from chat:", client.ChatID)
				delete(h.Clients[client.ChatID], client)
				close(client.Send)
			}
		case message := <-h.BroadCast:
			for client := range h.Clients[message.ChatID] {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients[message.ChatID], client)
				}
			}
		}
	}
}
