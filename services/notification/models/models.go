package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)

type NotifMessage struct {
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
}

type DeleteNotifMessage struct {
	NotifID int `json:"id"`
}

type NotifResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

type Signal struct {
	Sig int `json:"signal"`
}

type Client struct {
	ID          int
	SendChannel chan Signal
	Conn        *websocket.Conn
}

type Hub struct {
	Connect    chan *Client
	Disconnect chan *Client
	Send       chan int
	Clients    map[int]*Client
}

type JwtClaims struct {
	jwt.RegisteredClaims
	UserName string
	UserID   int
}
