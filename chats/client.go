package chats

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

type Client struct {
	hub    *Hub
	conn   *websocket.Conn
	send   chan *WebSocketMessage
	userID int
}

func NewClient(hub *Hub, conn *websocket.Conn, userID int) *Client {
	return &Client{
		hub:    hub,
		conn:   conn,
		send:   make(chan *WebSocketMessage, 256),
		userID: userID,
	}
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, msgData, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket read error: %v", err)
			}
			break
		}

		var message WebSocketMessage
		if err := json.Unmarshal(msgData, &message); err != nil {
			log.Printf("JSON parse error: %v", err)
			continue
		}

		message.UserID = c.userID
		message.Timestamp = time.Now().Unix()

		switch message.Type {
		case "message":
			// Сохраняем сообщение в БД и отправляем в комнату
			c.handleNewMessage(&message)

		case "typing":
			// Отправляем статус печатания
			c.hub.SendToRoom(message.ChatID, &message)

		case "read":
			// Отмечаем сообщения как прочитанные
			c.handleReadReceipt(&message)

		case "join":
			c.hub.JoinRoom(message.ChatID, c.userID)
			c.hub.SendToRoom(message.ChatID, &message)

		case "leave":
			c.hub.LeaveRoom(message.ChatID, c.userID)
			c.hub.SendToRoom(message.ChatID, &message)
		}
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			data, err := json.Marshal(message)
			if err != nil {
				continue
			}

			if err := c.conn.WriteMessage(websocket.TextMessage, data); err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) handleNewMessage(message *WebSocketMessage) {
	// Сохраняем сообщение в БД
	var msgData struct {
		Text    string `json:"text"`
		FileURL string `json:"file_url,omitempty"`
	}
	if err := json.Unmarshal(message.Message, &msgData); err != nil {
		return
	}

	// Здесь нужно вызвать метод сервиса для сохранения сообщения
	// Для упрощения отправляем сразу

	// Рассылаем сообщение всем в комнате
	c.hub.SendToRoom(message.ChatID, message)
}

func (c *Client) handleReadReceipt(message *WebSocketMessage) {
	// Отмечаем сообщения как прочитанные
	// Здесь нужно вызвать метод сервиса
	c.hub.SendToRoom(message.ChatID, message)
}
