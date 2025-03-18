package websocket

import (
	"log"

	"github.com/gorilla/websocket"
)

// Client represents a connected user
type Client struct {
	Conn *websocket.Conn
	Send chan []byte
}

// ReadMessages listens for messages from WebSocket
func (c *Client) ReadMessages(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}

		// Broadcast message to hub
		hub.Broadcast <- msg
	}
}

// WriteMessages sends messages to WebSocket client
func (c *Client) WriteMessages() {
	defer c.Conn.Close()
	for msg := range c.Send {
		c.Conn.WriteMessage(websocket.TextMessage, msg)
	}
}
