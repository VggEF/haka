package chats

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // В продакшене нужно настроить правильно
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WebSocketMessage struct {
	Type      string          `json:"type"` // message, typing, read, join, leave
	ChatID    int             `json:"chat_id"`
	UserID    int             `json:"user_id"`
	UserName  string          `json:"user_name,omitempty"`
	Message   json.RawMessage `json:"message,omitempty"`
	Timestamp int64           `json:"timestamp"`
}

type WebSocketHandler struct {
	hub *Hub
}

func NewWebSocketHandler(hub *Hub) *WebSocketHandler {
	return &WebSocketHandler{hub: hub}
}

func (h *WebSocketHandler) HandleWebSocket(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "не авторизован"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	client := NewClient(h.hub, conn, userID.(int))
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}
