package services

import (
	"sync"

	"github.com/ViitoJooj/clown-crm/internal/domain"
	"github.com/ViitoJooj/clown-crm/internal/repository"
	"github.com/gorilla/websocket"
)

type ChatHub struct {
	mu      sync.Mutex
	clients map[string]*websocket.Conn
	repo    repository.ChatRepository
}

func NewChatHub(repo repository.ChatRepository) *ChatHub {
	return &ChatHub{
		clients: make(map[string]*websocket.Conn),
		repo:    repo,
	}
}

func (h *ChatHub) AddClient(username string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[username] = conn
}

func (h *ChatHub) RemoveClient(username string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.clients, username)
}

func (h *ChatHub) SendTo(username string, msg domain.Chat) {
	h.mu.Lock()
	conn, ok := h.clients[username]
	h.mu.Unlock()

	if ok {
		_ = conn.WriteJSON(msg)
		_ = h.repo.SaveMessage(msg)
	}
}

func (h *ChatHub) Broadcast(msg domain.Chat) {
	h.mu.Lock()
	defer h.mu.Unlock()

	for username, conn := range h.clients {
		if err := conn.WriteJSON(msg); err != nil {
			delete(h.clients, username)
			conn.Close()
		}
	}

	_ = h.repo.SaveMessage(msg)
}
