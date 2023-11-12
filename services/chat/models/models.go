package models

import (
	"github.com/gorilla/websocket"
)

type Message struct {
	Message  string `json:"message"`
	UserID   int    `json:"user_id"`
	UserName string `json:"username"`
	Time     string `json:"time"`
	ChatID   string `json:"chat_id"`
}

type Client struct {
	Hub    *Hub
	Conn   *websocket.Conn
	Send   chan Message
	ChatID string
}

type Hub struct {
	Clients    map[string]map[*Client]bool
	BroadCast  chan Message
	Connect    chan *Client
	Disconnect chan *Client
}
