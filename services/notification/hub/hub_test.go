package hub

import (
	"testing"
	"time"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/models"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	New()
}


func TestUp(t *testing.T) {
	// Create a test hub
	hub := &models.Hub{
		Connect:    make(chan *models.Client),
		Disconnect: make(chan *models.Client),
		Send:       make(chan int),
		Clients:    make(map[int]*models.Client),
	}

	// Run the Up function in a goroutine
	go Up(hub)

	// Simulate a client connection
	clientID := 1
	client := &models.Client{ID: clientID}
	hub.Connect <- client

	// Allow some time for the goroutine to process the connection
	time.Sleep(time.Millisecond * 100)

	// Verify that the client is added to the hub's Clients map
	assert.Equal(t, client, hub.Clients[clientID], "Expected client to be connected to the hub")

	// Simulate a client disconnection
	hub.Disconnect <- client

	// Allow some time for the goroutine to process the disconnection
	time.Sleep(time.Millisecond * 100)

	// Verify that the client is removed from the hub's Clients map
	_, ok := hub.Clients[clientID]
	assert.False(t, ok, "Expected client to be disconnected from the hub")

	// Simulate sending a message to a client
	hub.Send <- clientID

	// Allow some time for the goroutine to process the message
	time.Sleep(time.Millisecond * 100)

	// Verify that the client receives the message
	// client.SendChannel <- models.Signal{Sig: 1}
	// select {
	// case receivedSignal := <-client.SendChannel:
	// 	assert.Equal(t, models.Signal{Sig: 1}, receivedSignal, "Expected client to receive the message")
	// default:
	// 	t.Error("Expected client to receive a message, but it did not happen")
	// }
}