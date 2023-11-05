package hub

import (
	"log"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/models"
)

func New() *models.Hub {
	return &models.Hub{
		Connect:    make(chan *models.Client),
		Disconnect: make(chan *models.Client),
		Clients:    make(map[int]*models.Client),
		Send:       make(chan int),
	}
}

func Up(h *models.Hub) {
	for {
		select {
		case msg := <-h.Connect:
			h.Clients[msg.ID] = msg
			log.Printf("User with id %d Connected\n", msg.ID)
		case msg := <-h.Disconnect:
			delete(h.Clients, msg.ID)
			log.Printf("User with id %d Disconnected\n", msg.ID)
		case msg := <-h.Send:
			_, ok := h.Clients[msg]
			if ok {
				h.Clients[msg].SendChannel <- models.Signal{Sig: 1}
			}
		}
	}
}
