package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/models"
	"github.com/gorilla/websocket"
)

func TestRead(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatalf("Failed to upgrade connection to WebSocket: %v", err)
		}
		defer conn.Close()

		client := &models.Client{Conn: conn}
		hub := &models.Hub{Disconnect: make(chan *models.Client)}
		go Read(client, hub)
		client.Conn.Close()
		<-hub.Disconnect
	}))
	client, _, err := websocket.DefaultDialer.Dial("ws"+server.URL[4:], nil)
	if err != nil {
		t.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	defer client.Close()
}


func TestWrite(t *testing.T) {
	// Create a mock WebSocket server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatalf("Failed to upgrade connection to WebSocket: %v", err)
		}
		defer conn.Close()

		client := &models.Client{Conn: conn, SendChannel: make(chan models.Signal)}
		// Run the Write function in a goroutine
		go Write(client)

		// Simulate sending a message
		message := models.Signal{Sig:1}
		client.SendChannel <- message
	}))

	defer server.Close()

	// Use a WebSocket client to connect to the mock server
	client, _, err := websocket.DefaultDialer.Dial("ws"+server.URL[4:], nil)
	if err != nil {
		t.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	defer client.Close()

	// Wait for the Write function to complete
	time.Sleep(time.Millisecond * 1000) // Allow some time for the Write function to complete
}