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
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatalf("Failed to upgrade connection to WebSocket: %v", err)
		}
		defer conn.Close()

		client := &models.Client{Conn: conn, SendChannel: make(chan models.Signal)}
		go Write(client)
		message := models.Signal{Sig: 1}
		client.SendChannel <- message
	}))

	defer server.Close()
	client, _, err := websocket.DefaultDialer.Dial("ws"+server.URL[4:], nil)
	if err != nil {
		t.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	defer client.Close()
	time.Sleep(time.Millisecond * 1000)
}
