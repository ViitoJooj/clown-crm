package domain

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Chat struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var BroadCast = make(chan Chat)

var Clients = make(map[string]*websocket.Conn)
