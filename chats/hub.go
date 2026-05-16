package chats

import (
	"log"
	"sync"
)

type Hub struct {
	clients    map[int]*Client      // userID -> client
	rooms      map[int]map[int]bool // chatID -> map[userID]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan *WebSocketMessage
	mu         sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[int]*Client),
		rooms:      make(map[int]map[int]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *WebSocketMessage, 256),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client.userID] = client
			h.mu.Unlock()
			log.Printf("Client %d connected. Total clients: %d", client.userID, len(h.clients))

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client.userID]; ok {
				delete(h.clients, client.userID)

				// Удаляем клиента из всех комнат
				for chatID, users := range h.rooms {
					delete(users, client.userID)
					if len(users) == 0 {
						delete(h.rooms, chatID)
					}
				}
			}
			h.mu.Unlock()
			close(client.send)
			log.Printf("Client %d disconnected. Total clients: %d", client.userID, len(h.clients))

		case message := <-h.broadcast:
			h.mu.RLock()
			// Отправляем всем в комнате
			if users, ok := h.rooms[message.ChatID]; ok {
				for userID := range users {
					if client, exists := h.clients[userID]; exists {
						select {
						case client.send <- message:
						default:
							close(client.send)
							delete(h.clients, client.userID)
						}
					}
				}
			}
			h.mu.RUnlock()
		}
	}
}

func (h *Hub) JoinRoom(chatID, userID int) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.rooms[chatID]; !ok {
		h.rooms[chatID] = make(map[int]bool)
	}
	h.rooms[chatID][userID] = true
}

func (h *Hub) LeaveRoom(chatID, userID int) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if users, ok := h.rooms[chatID]; ok {
		delete(users, userID)
		if len(users) == 0 {
			delete(h.rooms, chatID)
		}
	}
}

func (h *Hub) SendToRoom(chatID int, message *WebSocketMessage) {
	h.broadcast <- message
}

func (h *Hub) SendToUser(userID int, message *WebSocketMessage) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if client, ok := h.clients[userID]; ok {
		select {
		case client.send <- message:
		default:
			close(client.send)
			delete(h.clients, client.userID)
		}
	}
}
