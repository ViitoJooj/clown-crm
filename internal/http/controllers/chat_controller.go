package controllers

import (
	"log"

	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/ViitoJooj/clown-crm/internal/services"
	"github.com/gin-gonic/gin"
)

type ChatController struct {
	chatHub *services.ChatHub
}

func NewChatController(chatHub *services.ChatHub) *ChatController {
	return &ChatController{chatHub: chatHub}
}

func (c *ChatController) HandleConnections(ctx *gin.Context) {
	ws, err := domain.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	username := ctx.Query("username")
	if username == "" {
		log.Println("username vazio")
		return
	}

	c.chatHub.AddClient(username, ws)

	for {
		var msg domain.Chat
		if err := ws.ReadJSON(&msg); err != nil {
			c.chatHub.RemoveClient(username)
			break
		}

		if msg.To != "" {
			c.chatHub.SendTo(msg.To, msg)
		} else {
			c.chatHub.Broadcast(msg)
		}
	}
}
