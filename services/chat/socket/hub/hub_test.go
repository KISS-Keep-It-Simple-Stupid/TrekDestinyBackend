package hub

import (
	"testing"
	"time"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/chat/models"
	"github.com/stretchr/testify/assert"
)

func TestNewHub(t *testing.T) {
	NewHub()
}

func TestStartHub(t *testing.T) {
	hub := NewHub()
	go StartHub(hub)
	client1 := &models.Client{ChatID: "chat1", Send: make(chan models.Message)}
	client2 := &models.Client{ChatID: "chat1", Send: make(chan models.Message)}
	hub.Connect <- client1
	hub.Connect <- client2
	time.Sleep(100 * time.Millisecond)
	assert.Len(t, hub.Clients["chat1"], 2, "Unexpected number of clients in the map")
	hub.Disconnect <- client1
	time.Sleep(100 * time.Millisecond)
	assert.Len(t, hub.Clients["chat1"], 1, "Unexpected number of clients in the map after disconnection")
	message := models.Message{ChatID: "chat1", Message: "Hello, clients!"}
	hub.BroadCast <- message
	time.Sleep(100 * time.Millisecond)

	select {
	case  <-client1.Send:
	default:
		t.Error("Expected message not received by client1")
	}

	select {
	case  <-client2.Send:
	default:
		t.Error("Expected message not received by client2")
	}
	hub.Disconnect <- client2
	time.Sleep(100 * time.Millisecond)
}