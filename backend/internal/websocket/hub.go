package websocket

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/google/uuid"
)

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	mutex      sync.RWMutex
}

type Client struct {
	hub    *Hub
	conn   *websocket.Conn
	send   chan []byte
	userID uuid.UUID
}

type Message struct {
	Type    string      `json:"type"`
	UserID  string      `json:"user_id,omitempty"`
	Payload interface{} `json:"payload"`
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mutex.Lock()
			h.clients[client] = true
			h.mutex.Unlock()
			log.Printf("Client connected: %v", client.userID)

		case client := <-h.unregister:
			h.mutex.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			h.mutex.Unlock()
			log.Printf("Client disconnected: %v", client.userID)

		case message := <-h.broadcast:
			h.mutex.RLock()
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
			h.mutex.RUnlock()
		}
	}
}

func (h *Hub) BroadcastToUser(userID uuid.UUID, messageType string, payload interface{}) {
	message := Message{
		Type:    messageType,
		UserID:  userID.String(),
		Payload: payload,
	}

	data, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshaling message: %v", err)
		return
	}

	h.mutex.RLock()
	for client := range h.clients {
		if client.userID == userID {
			select {
			case client.send <- data:
			default:
				close(client.send)
				delete(h.clients, client)
			}
		}
	}
	h.mutex.RUnlock()
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			continue
		}

		// Handle different message types here
		switch msg.Type {
		case "ping":
			c.send <- []byte(`{"type":"pong"}`)
		}
	}
}

func (c *Client) writePump() {
	defer c.conn.Close()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}
		}
	}
}

func (h *Hub) HandleWebSocket(c *websocket.Conn, userID uuid.UUID) {
	client := &Client{
		hub:    h,
		conn:   c,
		send:   make(chan []byte, 256),
		userID: userID,
	}

	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}