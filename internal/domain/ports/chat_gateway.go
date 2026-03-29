package ports

import (
	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/gorilla/websocket"
)

type ChatGateway interface {
	AddClient(username string, conn *websocket.Conn)
	RemoveClient(username string)
	SendTo(username string, msg domain.Chat)
	Broadcast(msg domain.Chat)
}
